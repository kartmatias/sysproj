package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kartmatias/sysproj/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
