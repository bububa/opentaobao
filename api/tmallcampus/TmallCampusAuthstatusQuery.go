package tmallcampus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcampus"
)

// TmallCampusAuthstatusQuery 学生认证状态查询
// tmall.campus.authstatus.query
//
// 学生认证状态查询
func TmallCampusAuthstatusQuery(ctx context.Context, clt *core.SDKClient, req *tmallcampus.TmallCampusAuthstatusQueryAPIRequest, resp *tmallcampus.TmallCampusAuthstatusQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
