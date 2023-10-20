package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMallitemcenterSupplierAbilityUpdate 门店服务能力授权接口
// tmall.mallitemcenter.supplier.ability.update
//
// 门店服务能力授权
func TmallMallitemcenterSupplierAbilityUpdate(clt *core.SDKClient, req *tmallservice.TmallMallitemcenterSupplierAbilityUpdateAPIRequest, session string) (*tmallservice.TmallMallitemcenterSupplierAbilityUpdateAPIResponse, error) {
	var resp tmallservice.TmallMallitemcenterSupplierAbilityUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
