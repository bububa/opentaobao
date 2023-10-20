package ticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

// AlitripTicketRuleUpload 【门票API2.0】景点门票规则维护接口
// alitrip.ticket.rule.upload
//
// 景点门票规则维护接口。该接口同时支持新发规则和编辑现有规则，如果out_rule_id下没有发布过规则，则系统将判断为新发一个规则，否则认为是编辑现有规则。
// 对于新发布规则的情况，有些参数是必填的，请仔细查看各字段说明。对于编辑的情况，除out_rule_id外都是可选，编辑情况支持增量更新（某个参数不传则使用该规则上原有值）
func AlitripTicketRuleUpload(clt *core.SDKClient, req *ticket.AlitripTicketRuleUploadAPIRequest, resp *ticket.AlitripTicketRuleUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
