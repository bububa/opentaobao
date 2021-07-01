package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuBuyerTagGetAPIRequest
判断买家是否有某些标 API请求
taobao.qianniu.buyer.tag.get

判断某个买家是否有某些标 */
type TaobaoQianniuBuyerTagGetAPIRequest struct {
	model.Params
	// 买家nick
	_buyerNick string
	// 支持的表，多个tag用英文逗号切割
	_tagList string
}

// New
