package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsMessageShorturlQuery 聚石塔短链信息查询
// taobao.jst.sms.message.shorturl.query
//
// 聚石塔短链信息查询
func TaobaoJstSmsMessageShorturlQuery(clt *core.SDKClient, req *jst.TaobaoJstSmsMessageShorturlQueryAPIRequest, session string) (*jst.TaobaoJstSmsMessageShorturlQueryAPIResponse, error) {
	var resp jst.TaobaoJstSmsMessageShorturlQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
