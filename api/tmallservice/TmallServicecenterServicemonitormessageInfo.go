package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicemonitormessageInfo 服务预警单查询接口
// tmall.servicecenter.servicemonitormessage.info
//
// 服务预警单查询接口,用于查询预警单最新状态
func TmallServicecenterServicemonitormessageInfo(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterServicemonitormessageInfoAPIRequest, resp *tmallservice.TmallServicecenterServicemonitormessageInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
