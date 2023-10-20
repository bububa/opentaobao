package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateModify 淘宝短信模板修改
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
func TaobaoJstSmsTemplateModify(clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateModifyAPIRequest, resp *jst.TaobaoJstSmsTemplateModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
