package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicestoreUpdate 修改门店信息
// tmall.servicecenter.servicestore.update
//
// 用于修改门店/网点信息。多个业务共用
func TmallServicecenterServicestoreUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreUpdateAPIRequest, resp *tmallservice.TmallServicecenterServicestoreUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
