package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyDomfareCompareAPIRequest 比价工具 API请求
// alitrip.policy.domfare.compare
//
// 比价工具
type AlitripPolicyDomfareCompareAPIRequest struct {
	model.Params
	// 查询参数
	_compareDomFareRequestDTO *CompareDomFareRequestDto
}

// NewAlitripPolicyDomfareCompareRequest 初始化AlitripPolicyDomfareCompareAPIRequest对象
func NewAlitripPolicyDomfareCompareRequest() *AlitripPolicyDomfareCompareAPIRequest {
	return &AlitripPolicyDomfareCompareAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPolicyDomfareCompareAPIRequest) Reset() {
	r._compareDomFareRequestDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyDomfareCompareAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.domfare.compare"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyDomfareCompareAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicyDomfareCompareAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompareDomFareRequestDTO is CompareDomFareRequestDTO Setter
// 查询参数
func (r *AlitripPolicyDomfareCompareAPIRequest) SetCompareDomFareRequestDTO(_compareDomFareRequestDTO *CompareDomFareRequestDto) error {
	r._compareDomFareRequestDTO = _compareDomFareRequestDTO
	r.Set("compare_dom_fare_request_d_t_o", _compareDomFareRequestDTO)
	return nil
}

// GetCompareDomFareRequestDTO CompareDomFareRequestDTO Getter
func (r AlitripPolicyDomfareCompareAPIRequest) GetCompareDomFareRequestDTO() *CompareDomFareRequestDto {
	return r._compareDomFareRequestDTO
}

var poolAlitripPolicyDomfareCompareAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPolicyDomfareCompareRequest()
	},
}

// GetAlitripPolicyDomfareCompareRequest 从 sync.Pool 获取 AlitripPolicyDomfareCompareAPIRequest
func GetAlitripPolicyDomfareCompareAPIRequest() *AlitripPolicyDomfareCompareAPIRequest {
	return poolAlitripPolicyDomfareCompareAPIRequest.Get().(*AlitripPolicyDomfareCompareAPIRequest)
}

// ReleaseAlitripPolicyDomfareCompareAPIRequest 将 AlitripPolicyDomfareCompareAPIRequest 放入 sync.Pool
func ReleaseAlitripPolicyDomfareCompareAPIRequest(v *AlitripPolicyDomfareCompareAPIRequest) {
	v.Reset()
	poolAlitripPolicyDomfareCompareAPIRequest.Put(v)
}
