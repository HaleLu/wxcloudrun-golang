package service

import (
	"errors"
	"fmt"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/util"
)

var MissionServiceImpl MissionService

type MissionService struct {
	SimpleService[model.Mission, uint64]
}

func (s *MissionService) Finish(openId string, id uint64, finishTime time.Time) error {
	user, err := dao.UserDaoImpl.GetByOpenId(openId)
	if err != nil {
		return err
	}

	mission, err := s.Get(id)
	if err != nil {
		return err
	}
	if mission == nil {
		return errors.New("no this mission")
	}

	// check
	canFinish, err := s.CanFinish(openId, mission, finishTime)
	if err != nil {
		return err
	}
	if !canFinish {
		return errors.New("任务已经完成啦~")
	}
	// add mission log
	err = dao.OpLogDaoImpl.Upsert(&model.OpLog{
		Type:       model.OpTypeFinishMission,
		OpenId:     user.OpenId,
		ObjectId:   mission.Id,
		Price:      mission.Price,
		Content:    fmt.Sprintf("完成任务%s，余额%d->%d", mission.Title, user.Amount, user.Amount+mission.Price),
		FinishTime: finishTime,
	})
	if err != nil {
		return err
	}

	// update amount
	user.Amount += mission.Price
	err = dao.UserDaoImpl.Upsert(user)
	if err != nil {
		return err
	}

	// update last finish time
	mission.LastFinishTime = &finishTime
	err = dao.MissionDaoImpl.Upsert(mission)
	if err != nil {
		return err
	}
	return nil
}

func (s *MissionService) CanFinish(openId string, mission *model.Mission, finishTime time.Time) (bool, error) {
	if mission == nil {
		return false, nil
	}
	if mission.LastFinishTime == nil {
		return true, nil
	}
	switch mission.MissionType {
	case model.MissionTypeOnce:
		return false, nil
	case model.MissionTypeEveryday:
		hasFinished, err := s.hasFinishedInDate(openId, mission, finishTime)
		if err != nil {
			return false, err
		}
		return !hasFinished, nil
	case model.MissionTypeWorkdays:
		if !util.IsWorkday(finishTime.Weekday()) {
			return false, nil
		}
		hasFinished, err := s.hasFinishedInDate(openId, mission, finishTime)
		if err != nil {
			return false, err
		}
		return !hasFinished, nil
	case model.MissionTypeWeekends:
		if !util.IsWeekend(finishTime.Weekday()) {
			return false, nil
		}
		hasFinished, err := s.hasFinishedInDate(openId, mission, finishTime)
		if err != nil {
			return false, err
		}
		return !hasFinished, nil
	}
	return false, nil
}

func (s *MissionService) hasFinishedInDate(openId string, mission *model.Mission, finishTime time.Time) (bool, error) {
	opLogs, err := dao.OpLogDaoImpl.GetByDate(openId, model.OpTypeFinishMission, mission.Id, finishTime)
	if err != nil {
		return false, err
	}
	return len(opLogs) > 0, nil
}
