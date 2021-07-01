package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* AlibabaJymIndustryInformationCallbak
VMOS回调行业信息系统
alibaba.jym.industry.information.callbak

VMOS回调交易猫行业信息系统 */
func AlibabaJymIndustryInformationCallbak(clt *core.SDKClient, req *product.AlibabaJymIndustryInformationCallbakAPIRequest, session string) (*product.AlibabaJymIndustryInformationCallbakAPIResponse, error) {
	var resp product.AlibabaJymIndustryInformationCallbakAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
