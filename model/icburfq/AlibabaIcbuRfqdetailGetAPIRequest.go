package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqdetailGetAPIRequest 获取RFQ详情 API请求
// alibaba.icbu.rfqdetail.get
//
// 查看RFQ的详情信息
type AlibabaIcbuRfqdetailGetAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 查询RFQ详情DTO
	_rfqQueryDto *RfqDetailSearchQueryDto
}

// NewAlibabaIcbuRfqdetailGetRequest 初始化AlibabaIcbuRfqdetailGetAPIRequest对象
func NewAlibabaIcbuRfqdetailGetRequest() *AlibabaIcbuRfqdetailGetAPIRequest {
	return &AlibabaIcbuRfqdetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqdetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfqdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqdetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMd5key is Md5key Setter
// 验证
func (r *AlibabaIcbuRfqdetailGetAPIRequest) SetMd5key(_md5key string) error {
	r._md5key = _md5key
	r.Set("md5key", _md5key)
	return nil
}

// GetMd5key Md5key Getter
func (r AlibabaIcbuRfqdetailGetAPIRequest) GetMd5key() string {
	return r._md5key
}

// SetRfqQueryDto is RfqQueryDto Setter
// 查询RFQ详情DTO
func (r *AlibabaIcbuRfqdetailGetAPIRequest) SetRfqQueryDto(_rfqQueryDto *RfqDetailSearchQueryDto) error {
	r._rfqQueryDto = _rfqQueryDto
	r.Set("rfq_query_dto", _rfqQueryDto)
	return nil
}

// GetRfqQueryDto RfqQueryDto Getter
func (r AlibabaIcbuRfqdetailGetAPIRequest) GetRfqQueryDto() *RfqDetailSearchQueryDto {
	return r._rfqQueryDto
}
