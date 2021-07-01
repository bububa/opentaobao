package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuQuotationPostAPIRequest
供应商提交报价 API请求
alibaba.icbu.quotation.post

供应商对RFQ进行报价 */
type AlibabaIcbuQuotationPostAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 报价DTO
	_dto *RfqQuotationRemoteDto
}

// New
