package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicydomfarecompareAPIRequest 比价工具 API请求
// alitrip.policy.domfare.compare
//
// 比价工具
type AlitrippolicydomfarecompareAPIRequest struct {
	model.Params
	// 查询参数
	_compareDomFareRequestDTO *CompareDomFareRequestDto
}

// NewAlitrippolicydomfarecompareRequest 初始化AlitrippolicydomfarecompareAPIRequest对象
func NewAlitrippolicydomfarecompareRequest() *AlitrippolicydomfarecompareAPIRequest {
	return &AlitrippolicydomfarecompareAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitrippolicydomfarecompareAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.domfare.compare"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitrippolicydomfarecompareAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitrippolicydomfarecompareAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompareDomFareRequestDTO is CompareDomFareRequestDTO Setter
// 查询参数
func (r *AlitrippolicydomfarecompareAPIRequest) SetCompareDomFareRequestDTO(_compareDomFareRequestDTO *CompareDomFareRequestDto) error {
	r._compareDomFareRequestDTO = _compareDomFareRequestDTO
	r.Set("compare_dom_fare_request_d_t_o", _compareDomFareRequestDTO)
	return nil
}

// GetCompareDomFareRequestDTO CompareDomFareRequestDTO Getter
func (r AlitrippolicydomfarecompareAPIRequest) GetCompareDomFareRequestDTO() *CompareDomFareRequestDto {
	return r._compareDomFareRequestDTO
}
