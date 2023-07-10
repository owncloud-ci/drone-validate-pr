package plugin

import (
	"context"

	"github.com/drone/drone-go/plugin/validator"
	"github.com/sirupsen/logrus"
)

// New returns a new validator plugin.
func New() validator.Plugin {
	return &plugin{}
}

type plugin struct{}

//nolint:revive
func (p *plugin) Validate(ctx context.Context, req *validator.Request) error {
	switch req.Build.Event {
	case "push", "tag": // triggered by folks with write access to the repo, therefore trusted
		logrus.Debugf("%s build ignored", req.Build.Event)

		return nil
	case "cron", "promote", "rollback", "custom": // triggered by folks with write access in drone
		logrus.Debugf("%s build ignored", req.Build.Event)

		return nil
	//nolint:goconst
	case "pull_request": // may be triggered by folks without write access, needs approval
		if !isFork(req) {
			return nil
		}

		fallthrough
	default: // unknown new event, fail secure
		logrus.Warnf("%s build unrecognized", req.Build.Event)

		return validator.ErrBlock
	}
}

func isFork(req *validator.Request) bool {
	sourceRepo := req.Build.Fork
	targetRepo := req.Repo.Slug

	if sourceRepo != targetRepo {
		logrus.WithFields(
			logrus.Fields{"source": sourceRepo, "target": targetRepo}).Infof("%s needs approval", req.Build.Link)
	} else {
		logrus.WithFields(logrus.Fields{"source": sourceRepo, "target": targetRepo}).Infof("%s approved", req.Build.Link)
	}

	return sourceRepo == targetRepo
}
