package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymIndustryInformationCallbak VMOS回调行业信息系统
// alibaba.jym.industry.information.callbak
//
// VMOS回调交易猫行业信息系统
func AlibabaJymIndustryInformationCallbak(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymIndustryInformationCallbakAPIRequest, resp *product.AlibabaJymIndustryInformationCallbakAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
