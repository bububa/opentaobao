package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuquotationpostAPIRequest 供应商提交报价 API请求
// alibaba.icbu.quotation.post
//
// 供应商对RFQ进行报价
type AlibabaicbuquotationpostAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 报价DTO
	_dto *RfqQuotationRemoteDto
}

// NewAlibabaicbuquotationpostRequest 初始化AlibabaicbuquotationpostAPIRequest对象
func NewAlibabaicbuquotationpostRequest() *AlibabaicbuquotationpostAPIRequest {
	return &AlibabaicbuquotationpostAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuquotationpostAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.quotation.post"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuquotationpostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuquotationpostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMd5key is Md5key Setter
// 验证
func (r *AlibabaicbuquotationpostAPIRequest) SetMd5key(_md5key string) error {
	r._md5key = _md5key
	r.Set("md5key", _md5key)
	return nil
}

// GetMd5key Md5key Getter
func (r AlibabaicbuquotationpostAPIRequest) GetMd5key() string {
	return r._md5key
}

// SetDto is Dto Setter
// 报价DTO
func (r *AlibabaicbuquotationpostAPIRequest) SetDto(_dto *RfqQuotationRemoteDto) error {
	r._dto = _dto
	r.Set("dto", _dto)
	return nil
}

// GetDto Dto Getter
func (r AlibabaicbuquotationpostAPIRequest) GetDto() *RfqQuotationRemoteDto {
	return r._dto
}
