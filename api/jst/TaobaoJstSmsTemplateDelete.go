package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateDelete 淘宝短信模板删除
// taobao.jst.sms.template.delete
//
// 淘宝短信模板删除
func TaobaoJstSmsTemplateDelete(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateDeleteAPIRequest, resp *jst.TaobaoJstSmsTemplateDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
