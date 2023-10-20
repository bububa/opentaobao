package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaWholesaleShippinglineTemplateList 获取运费模板
// alibaba.wholesale.shippingline.template.list
//
// 查询运费模板信息
func AlibabaWholesaleShippinglineTemplateList(clt *core.SDKClient, req *icbu.AlibabaWholesaleShippinglineTemplateListAPIRequest, resp *icbu.AlibabaWholesaleShippinglineTemplateListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
