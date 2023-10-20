package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyProductGet 备孕首页获取达人配置商品
// alibaba.alihealth.pregnancy.product.get
//
// 备孕首页获取达人配置商品
func AlibabaAlihealthPregnancyProductGet(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyProductGetAPIRequest, resp *alihealthcrm.AlibabaAlihealthPregnancyProductGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
