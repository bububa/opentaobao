package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsMessageShorturlCreate 聚石塔营销效果短链生成
// taobao.jst.sms.message.shorturl.create
//
// 聚石塔生成淘短链接口
func TaobaoJstSmsMessageShorturlCreate(clt *core.SDKClient, req *jst.TaobaoJstSmsMessageShorturlCreateAPIRequest, resp *jst.TaobaoJstSmsMessageShorturlCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
