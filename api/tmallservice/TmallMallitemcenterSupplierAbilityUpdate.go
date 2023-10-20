package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMallitemcenterSupplierAbilityUpdate 门店服务能力授权接口
// tmall.mallitemcenter.supplier.ability.update
//
// 门店服务能力授权
func TmallMallitemcenterSupplierAbilityUpdate(clt *core.SDKClient, req *tmallservice.TmallMallitemcenterSupplierAbilityUpdateAPIRequest, resp *tmallservice.TmallMallitemcenterSupplierAbilityUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
