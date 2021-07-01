package newretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItApAddressGetAPIRequest
getApAddressByMacNew API请求
alibaba.it.ap.address.get

根据ap 的mac地址查询ap的结构化位置信息 */
type AlibabaItApAddressGetAPIRequest struct {
	model.Params
	// 签名
	_signature string
	// 语言
	_language string
	// 分配的ak
	_appKeyInternal string
	// ap的mac
	_mac string
	// 当前时间戳，毫秒
	_timestampInternal int64
}

// New
