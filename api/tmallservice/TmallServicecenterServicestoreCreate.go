package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicestoreCreate 创建门店
// tmall.servicecenter.servicestore.create
//
// 用于创建门店/网点。多个业务共用
func TmallServicecenterServicestoreCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreCreateAPIRequest, resp *tmallservice.TmallServicecenterServicestoreCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
