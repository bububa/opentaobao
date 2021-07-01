package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMessageShorturlQueryAPIRequest
聚石塔短链信息查询 API请求
taobao.jst.sms.message.shorturl.query

聚石塔短链信息查询 */
type TaobaoJstSmsMessageShorturlQueryAPIRequest struct {
	model.Params
	// 短链名，即域名后的字符串
	_shortName string
}

// New
