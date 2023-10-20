package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateDelete 淘宝短信模板删除
// taobao.jst.sms.template.delete
//
// 淘宝短信模板删除
func TaobaoJstSmsTemplateDelete(clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateDeleteAPIRequest, session string) (*jst.TaobaoJstSmsTemplateDeleteAPIResponse, error) {
	var resp jst.TaobaoJstSmsTemplateDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
