package plugin

import (
	"context"
	"errors"
	"testing"

	"github.com/drone/drone-go/plugin/validator"
)

//nolint:goconst
func TestForkPullRequest(t *testing.T) {
	p := &plugin{}
	req := validator.Request{}
	req.Build.Event = "pull_request"
	req.Build.Fork = "valid/drone-test"
	req.Repo.Slug = "fork/drone-test"

	err := p.Validate(context.Background(), &req)
	if !errors.Is(err, validator.ErrBlock) {
		t.Fatal("expected PR from fork to be blocked")
	}
}

func TestSameRepoPullRequest(t *testing.T) {
	p := &plugin{}
	req := validator.Request{}
	req.Build.Event = "pull_request"
	req.Build.Fork = "fork/drone-test"
	req.Repo.Slug = "fork/drone-test"

	err := p.Validate(context.Background(), &req)
	if err != nil {
		t.Fatal("expected PR from the same repo to be approved")
	}
}

func TestUnknownEvent(t *testing.T) {
	p := &plugin{}
	req := validator.Request{}
	req.Build.Event = "merge_request"

	err := p.Validate(context.Background(), &req)
	if !errors.Is(err, validator.ErrBlock) {
		t.Fatal("expected build with unknown Event to be blocked")
	}
}
