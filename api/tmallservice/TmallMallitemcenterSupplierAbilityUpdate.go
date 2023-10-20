package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallmallitemcentersupplierabilityupdate 门店服务能力授权接口
// tmall.mallitemcenter.supplier.ability.update
//
// 门店服务能力授权
func Tmallmallitemcentersupplierabilityupdate(clt *core.SDKClient, req *tmallservice.TmallmallitemcentersupplierabilityupdateAPIRequest, session string) (*tmallservice.TmallmallitemcentersupplierabilityupdateAPIResponse, error) {
	var resp tmallservice.TmallmallitemcentersupplierabilityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
