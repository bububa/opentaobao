package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordStatusUpdateAPIRequest
关键词启动暂停推广 API请求
alibaba.scbp.ad.keyword.status.update

关键词启动暂停推广 */
type AlibabaScbpAdKeywordStatusUpdateAPIRequest struct {
	model.Params
	// 只能取ascci字符
	_adKeyword string
	// 只能去in_promotion或者stopped
	_status string
}

// New
