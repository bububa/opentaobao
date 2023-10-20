package icburfq

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuRfqdetailGetAPIRequest) Reset() {
	r._md5key = ""
	r._rfqQueryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqdetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfqdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqdetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuRfqdetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIcbuRfqdetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuRfqdetailGetRequest()
	},
}

// GetAlibabaIcbuRfqdetailGetRequest 从 sync.Pool 获取 AlibabaIcbuRfqdetailGetAPIRequest
func GetAlibabaIcbuRfqdetailGetAPIRequest() *AlibabaIcbuRfqdetailGetAPIRequest {
	return poolAlibabaIcbuRfqdetailGetAPIRequest.Get().(*AlibabaIcbuRfqdetailGetAPIRequest)
}

// ReleaseAlibabaIcbuRfqdetailGetAPIRequest 将 AlibabaIcbuRfqdetailGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuRfqdetailGetAPIRequest(v *AlibabaIcbuRfqdetailGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuRfqdetailGetAPIRequest.Put(v)
}
