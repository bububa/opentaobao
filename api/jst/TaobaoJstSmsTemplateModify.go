package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateModify 淘宝短信模板修改
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
func TaobaoJstSmsTemplateModify(clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateModifyAPIRequest, session string) (*jst.TaobaoJstSmsTemplateModifyAPIResponse, error) {
	var resp jst.TaobaoJstSmsTemplateModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
