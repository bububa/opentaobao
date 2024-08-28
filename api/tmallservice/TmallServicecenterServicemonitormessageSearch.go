package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicemonitormessageSearch 根据时间段查询服务商的服务预警消息列表(15分钟内)
// tmall.servicecenter.servicemonitormessage.search
//
// 根据时间段查询服务商的服务预警消息列表(15分钟内)
func TmallServicecenterServicemonitormessageSearch(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterServicemonitormessageSearchAPIRequest, resp *tmallservice.TmallServicecenterServicemonitormessageSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
