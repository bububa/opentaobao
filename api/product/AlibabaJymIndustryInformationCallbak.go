package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymindustryinformationcallbak VMOS回调行业信息系统
// alibaba.jym.industry.information.callbak
//
// VMOS回调交易猫行业信息系统
func Alibabajymindustryinformationcallbak(clt *core.SDKClient, req *product.AlibabajymindustryinformationcallbakAPIRequest, session string) (*product.AlibabajymindustryinformationcallbakAPIResponse, error) {
	var resp product.AlibabajymindustryinformationcallbakAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
