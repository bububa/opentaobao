package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateCreate 淘宝短信模板创建
// taobao.jst.sms.template.create
//
// 聚石塔短信模板创建
func TaobaoJstSmsTemplateCreate(clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateCreateAPIRequest, resp *jst.TaobaoJstSmsTemplateCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
