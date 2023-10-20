package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateQuery 淘宝短信模板查询
// taobao.jst.sms.template.query
//
// 淘宝短信模板查询
func TaobaoJstSmsTemplateQuery(clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateQueryAPIRequest, resp *jst.TaobaoJstSmsTemplateQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
