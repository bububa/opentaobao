package icburfq

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqRecommendAPIRequest rfq推荐 API请求
// alibaba.icbu.rfq.recommend
//
// rfq推荐
type AlibabaIcbuRfqRecommendAPIRequest struct {
	model.Params
	// 入参数据
	_queryDto *QueryDto
}

// NewAlibabaIcbuRfqRecommendRequest 初始化AlibabaIcbuRfqRecommendAPIRequest对象
func NewAlibabaIcbuRfqRecommendRequest() *AlibabaIcbuRfqRecommendAPIRequest {
	return &AlibabaIcbuRfqRecommendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuRfqRecommendAPIRequest) Reset() {
	r._queryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqRecommendAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqRecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuRfqRecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// 入参数据
func (r *AlibabaIcbuRfqRecommendAPIRequest) SetQueryDto(_queryDto *QueryDto) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaIcbuRfqRecommendAPIRequest) GetQueryDto() *QueryDto {
	return r._queryDto
}

var poolAlibabaIcbuRfqRecommendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuRfqRecommendRequest()
	},
}

// GetAlibabaIcbuRfqRecommendRequest 从 sync.Pool 获取 AlibabaIcbuRfqRecommendAPIRequest
func GetAlibabaIcbuRfqRecommendAPIRequest() *AlibabaIcbuRfqRecommendAPIRequest {
	return poolAlibabaIcbuRfqRecommendAPIRequest.Get().(*AlibabaIcbuRfqRecommendAPIRequest)
}

// ReleaseAlibabaIcbuRfqRecommendAPIRequest 将 AlibabaIcbuRfqRecommendAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuRfqRecommendAPIRequest(v *AlibabaIcbuRfqRecommendAPIRequest) {
	v.Reset()
	poolAlibabaIcbuRfqRecommendAPIRequest.Put(v)
}
