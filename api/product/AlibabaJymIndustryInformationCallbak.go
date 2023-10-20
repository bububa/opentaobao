package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymIndustryInformationCallbak VMOS回调行业信息系统
// alibaba.jym.industry.information.callbak
//
// VMOS回调交易猫行业信息系统
func AlibabaJymIndustryInformationCallbak(clt *core.SDKClient, req *product.AlibabaJymIndustryInformationCallbakAPIRequest, resp *product.AlibabaJymIndustryInformationCallbakAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
