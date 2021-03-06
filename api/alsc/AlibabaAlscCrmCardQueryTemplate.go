package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardQueryTemplate 查询卡模板详情
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
func AlibabaAlscCrmCardQueryTemplate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardQueryTemplateAPIRequest, session string) (*alsc.AlibabaAlscCrmCardQueryTemplateAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmCardQueryTemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
