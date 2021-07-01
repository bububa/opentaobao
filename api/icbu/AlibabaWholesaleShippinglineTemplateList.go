package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

/* AlibabaWholesaleShippinglineTemplateList
获取运费模板
alibaba.wholesale.shippingline.template.list

查询运费模板信息 */
func AlibabaWholesaleShippinglineTemplateList(clt *core.SDKClient, req *icbu.AlibabaWholesaleShippinglineTemplateListAPIRequest, session string) (*icbu.AlibabaWholesaleShippinglineTemplateListAPIResponse, error) {
	var resp icbu.AlibabaWholesaleShippinglineTemplateListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
