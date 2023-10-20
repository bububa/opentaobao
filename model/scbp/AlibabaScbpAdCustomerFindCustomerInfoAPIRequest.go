package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCustomerFindCustomerInfoAPIRequest 查询客户信息 API请求
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
type AlibabaScbpAdCustomerFindCustomerInfoAPIRequest struct {
	model.Params
	// 用户信息
	_topContextDto *TopContextDto
}

// NewAlibabaScbpAdCustomerFindCustomerInfoRequest 初始化AlibabaScbpAdCustomerFindCustomerInfoAPIRequest对象
func NewAlibabaScbpAdCustomerFindCustomerInfoRequest() *AlibabaScbpAdCustomerFindCustomerInfoAPIRequest {
	return &AlibabaScbpAdCustomerFindCustomerInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) Reset() {
	r._topContextDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.customer.find.customer.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContextDto is TopContextDto Setter
// 用户信息
func (r *AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) SetTopContextDto(_topContextDto *TopContextDto) error {
	r._topContextDto = _topContextDto
	r.Set("top_context_dto", _topContextDto)
	return nil
}

// GetTopContextDto TopContextDto Getter
func (r AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) GetTopContextDto() *TopContextDto {
	return r._topContextDto
}

var poolAlibabaScbpAdCustomerFindCustomerInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCustomerFindCustomerInfoRequest()
	},
}

// GetAlibabaScbpAdCustomerFindCustomerInfoRequest 从 sync.Pool 获取 AlibabaScbpAdCustomerFindCustomerInfoAPIRequest
func GetAlibabaScbpAdCustomerFindCustomerInfoAPIRequest() *AlibabaScbpAdCustomerFindCustomerInfoAPIRequest {
	return poolAlibabaScbpAdCustomerFindCustomerInfoAPIRequest.Get().(*AlibabaScbpAdCustomerFindCustomerInfoAPIRequest)
}

// ReleaseAlibabaScbpAdCustomerFindCustomerInfoAPIRequest 将 AlibabaScbpAdCustomerFindCustomerInfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCustomerFindCustomerInfoAPIRequest(v *AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCustomerFindCustomerInfoAPIRequest.Put(v)
}
