package api

import (
	"fmt"
	"go.uber.org/zap"
)

func (s *Service) GetSeed() error {
	resp, err := s.Client.R().
		Post("/seed/claim")
	if err != nil {
		return fmt.Errorf("error claiming seed: %v", err)
	}

	if resp.StatusCode() == 200 {
		s.Logger.Info("✅ SEED claimed successfully!")
	}
	return nil
}

func (s *Service) BuyUpgrade() error {
	_, err := s.Client.R().Post("/seed/mining-speed/upgrade")
	if err != nil {
		s.Logger.Error("Error while buying upgrade:", zap.Error(err))
		return err
	}
	return nil
}
