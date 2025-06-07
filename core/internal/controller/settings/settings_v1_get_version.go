package settings

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/settings/v1"
)

func (c *ControllerV1) GetVersion(ctx context.Context, req *v1.GetVersionReq) (res *v1.GetVersionRes, err error) {
	res = &v1.GetVersionRes{}

	res.Data.Version = g.Cfg().MustGet(ctx, "server.version", "0.1").String()

	res.SetSuccess("Success")

	return
}
