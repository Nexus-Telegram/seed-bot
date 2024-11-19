package api

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"nexus-seed-bot/types"
)

type CompleteUpgradeTaskPayload struct {
	Data struct {
		Id   string `json:"id"`
		Data struct {
			Id      string `json:"id"`
			Repeats int    `json:"repeats"`
		} `json:"data"`
	} `json:"data"`
}

//func (s *Service) CompleteUpgradeTask(taskID string) error {
//	payload := CompleteUpgradeTaskPayload{
//		Data: {
//			Id: taskID,
//		},
//	}
//	res, err := s.Client.R().SetBody(payload).Post(fmt.Sprintf("/tasks/%s", taskID))
//	if err != nil {
//		s.Logger.Error(err.Error())
//		return err
//	}
//	if res.StatusCode() == 200 {
//		s.Logger.Info(fmt.Sprintf("Success completing task: %s", taskID))
//	}
//	return nil
//}

func (s *Service) CompleteTaskWithoutConfirmation(taskID string) error {
	payload := "{}"
	res, err := s.Client.R().SetBody(payload).Post(fmt.Sprintf("/tasks/%s", taskID))
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	if res.StatusCode() == 200 {
		s.Logger.Info(fmt.Sprintf("Success completing task: %s", taskID))
	}
	return nil
}

type TaskCompleted struct {
	Id string `json:"data"`
}

func (s *Service) CompleteTaskWithSecret(taskID string, answer string) (*TaskCompleted, error) {
	payload := types.TaskAnswerPayload{
		Answer: answer,
	}
	var taskCompletedId TaskCompleted

	res, err := s.Client.R().SetResult(&taskCompletedId).SetBody(payload).Post(fmt.Sprintf("/tasks/%s", taskID))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		s.Logger.Error("Failed to fetch daily login streak",
			zap.String("status", res.Status()),
			zap.Any("data", string(res.Body())),
		)
		return nil, errors.New(fmt.Sprintf("Failed to complete task %s", taskID))
	}

	return &taskCompletedId, nil
}

func (s *Service) GetProgress() *types.Progress {
	res, err := s.Client.R().SetResult(&types.Progress{}).Get("/tasks/progresses")
	if err != nil {
		s.Logger.Error(err.Error())
		return nil
	}
	return res.Result().(*types.Progress)
}
