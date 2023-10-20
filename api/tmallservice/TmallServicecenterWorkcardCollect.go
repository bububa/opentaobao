package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardCollect 工单揽件
// tmall.servicecenter.workcard.collect
//
// 服务工单揽件接口
func TmallServicecenterWorkcardCollect(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCollectAPIRequest, resp *tmallservice.TmallServicecenterWorkcardCollectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
