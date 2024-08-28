package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateModify 淘宝短信模板修改
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
func TaobaoJstSmsTemplateModify(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateModifyAPIRequest, resp *jst.TaobaoJstSmsTemplateModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
