package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardQueryTemplate 查询卡模板详情
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
func AlibabaAlscCrmCardQueryTemplate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardQueryTemplateAPIRequest, resp *alsc.AlibabaAlscCrmCardQueryTemplateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
