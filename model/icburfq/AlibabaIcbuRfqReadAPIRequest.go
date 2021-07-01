package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuRfqReadAPIRequest
是否已读RFQ API请求
alibaba.icbu.rfq.read

是否已读RFQ */
type AlibabaIcbuRfqReadAPIRequest struct {
	model.Params
	// 查询RFQID列表
	_rfqIdList []string
}

// New
