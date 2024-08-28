package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsMessageShorturlQuery 聚石塔短链信息查询
// taobao.jst.sms.message.shorturl.query
//
// 聚石塔短链信息查询
func TaobaoJstSmsMessageShorturlQuery(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsMessageShorturlQueryAPIRequest, resp *jst.TaobaoJstSmsMessageShorturlQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
