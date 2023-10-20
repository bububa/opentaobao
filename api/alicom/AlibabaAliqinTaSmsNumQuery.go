package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinTaSmsNumQuery 短信查询
// alibaba.aliqin.ta.sms.num.query
//
// 查询短信发送揭露
func AlibabaAliqinTaSmsNumQuery(clt *core.SDKClient, req *alicom.AlibabaAliqinTaSmsNumQueryAPIRequest, resp *alicom.AlibabaAliqinTaSmsNumQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
