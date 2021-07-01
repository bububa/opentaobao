package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMessageShorturlCreateAPIRequest
聚石塔营销效果短链生成 API请求
taobao.jst.sms.message.shorturl.create

聚石塔生成淘短链接口 */
type TaobaoJstSmsMessageShorturlCreateAPIRequest struct {
	model.Params
	// 是否需要https前缀： true-要  false-不要
	_needHttpsPrefix bool
	// 人群标签
	_tag string
	// 商品或者店铺的H5地址，只支持长链
	_url string
	// 批次号
	_batchNumber string
}

// New
