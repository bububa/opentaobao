package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicburfqdetailgetAPIRequest 获取RFQ详情 API请求
// alibaba.icbu.rfqdetail.get
//
// 查看RFQ的详情信息
type AlibabaicburfqdetailgetAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 查询RFQ详情DTO
	_rfqQueryDto *RfqDetailSearchQueryDto
}

// NewAlibabaicburfqdetailgetRequest 初始化AlibabaicburfqdetailgetAPIRequest对象
func NewAlibabaicburfqdetailgetRequest() *AlibabaicburfqdetailgetAPIRequest {
	return &AlibabaicburfqdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicburfqdetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfqdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicburfqdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicburfqdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMd5key is Md5key Setter
// 验证
func (r *AlibabaicburfqdetailgetAPIRequest) SetMd5key(_md5key string) error {
	r._md5key = _md5key
	r.Set("md5key", _md5key)
	return nil
}

// GetMd5key Md5key Getter
func (r AlibabaicburfqdetailgetAPIRequest) GetMd5key() string {
	return r._md5key
}

// SetRfqQueryDto is RfqQueryDto Setter
// 查询RFQ详情DTO
func (r *AlibabaicburfqdetailgetAPIRequest) SetRfqQueryDto(_rfqQueryDto *RfqDetailSearchQueryDto) error {
	r._rfqQueryDto = _rfqQueryDto
	r.Set("rfq_query_dto", _rfqQueryDto)
	return nil
}

// GetRfqQueryDto RfqQueryDto Getter
func (r AlibabaicburfqdetailgetAPIRequest) GetRfqQueryDto() *RfqDetailSearchQueryDto {
	return r._rfqQueryDto
}
