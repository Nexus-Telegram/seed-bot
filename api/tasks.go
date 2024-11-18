package api

import (
	"fmt"
	"nexus-seed-bot/types"
)

func (s *Service) CompleteTask(taskID string) error {
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
func (s *Service) GetProgress() *types.Progress {
	res, err := s.Client.R().SetResult(&types.Progress{}).Get("/tasks/progresses")
	if err != nil {
		s.Logger.Error(err.Error())
		return nil
	}
	return res.Result().(*types.Progress)
}
