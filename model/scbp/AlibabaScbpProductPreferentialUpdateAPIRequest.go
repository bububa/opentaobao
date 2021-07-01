package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpProductPreferentialUpdateAPIRequest
设置P4P产品优先推广状态 API请求
alibaba.scbp.product.preferential.update

设置P4P产品优先推广状态 */
type AlibabaScbpProductPreferentialUpdateAPIRequest struct {
	model.Params
	// 关键词ID
	_keywordId int64
	// 产品ID
	_productId int64
	// Y:设置优推,N:取消优推
	_status string
}

// New
