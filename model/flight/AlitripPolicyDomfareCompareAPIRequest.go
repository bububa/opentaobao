package flight

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
