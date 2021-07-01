package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuRfqSearchAPIRequest
查询RFQ API请求
alibaba.icbu.rfq.search

用于查询RFQ的信息 */
type AlibabaIcbuRfqSearchAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 查询条件
	_cond *RfqRequestSearchCondDto
}

// New
