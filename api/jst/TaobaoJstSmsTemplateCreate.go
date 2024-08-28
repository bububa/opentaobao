package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateCreate 淘宝短信模板创建
// taobao.jst.sms.template.create
//
// 聚石塔短信模板创建
func TaobaoJstSmsTemplateCreate(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateCreateAPIRequest, resp *jst.TaobaoJstSmsTemplateCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
