package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoChainstoreRangesAPIRequest
蜂鸟查询门店配送范围接口 API请求
alibaba.ele.fengniao.chainstore.ranges

蜂鸟查询门店配送范围接口 */
type AlibabaEleFengniaoChainstoreRangesAPIRequest struct {
	model.Params
	// 商户code
	_merchantCode string
	// appId
	_appId string
	// 门店code
	_chainstoreCode string
}

// New
