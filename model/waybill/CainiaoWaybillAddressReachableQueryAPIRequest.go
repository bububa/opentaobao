package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillAddressReachableQueryAPIRequest 地址可达查询 API请求
// cainiao.waybill.address.reachable.query
//
// 地址可达查询
type CainiaoWaybillAddressReachableQueryAPIRequest struct {
	model.Params
	// 入参
	_reachableRecommendRequestDto *ReachableRecommendRequestDto
	// 调用方对象
	_clientInfoDto *ClientInfoDto
}

// NewCainiaoWaybillAddressReachableQueryRequest 初始化CainiaoWaybillAddressReachableQueryAPIRequest对象
func NewCainiaoWaybillAddressReachableQueryRequest() *CainiaoWaybillAddressReachableQueryAPIRequest {
	return &CainiaoWaybillAddressReachableQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillAddressReachableQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.address.reachable.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillAddressReachableQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReachableRecommendRequestDto is ReachableRecommendRequestDto Setter
// 入参
func (r *CainiaoWaybillAddressReachableQueryAPIRequest) SetReachableRecommendRequestDto(_reachableRecommendRequestDto *ReachableRecommendRequestDto) error {
	r._reachableRecommendRequestDto = _reachableRecommendRequestDto
	r.Set("reachable_recommend_request_dto", _reachableRecommendRequestDto)
	return nil
}

// GetReachableRecommendRequestDto ReachableRecommendRequestDto Getter
func (r CainiaoWaybillAddressReachableQueryAPIRequest) GetReachableRecommendRequestDto() *ReachableRecommendRequestDto {
	return r._reachableRecommendRequestDto
}

// SetClientInfoDto is ClientInfoDto Setter
// 调用方对象
func (r *CainiaoWaybillAddressReachableQueryAPIRequest) SetClientInfoDto(_clientInfoDto *ClientInfoDto) error {
	r._clientInfoDto = _clientInfoDto
	r.Set("client_info_dto", _clientInfoDto)
	return nil
}

// GetClientInfoDto ClientInfoDto Getter
func (r CainiaoWaybillAddressReachableQueryAPIRequest) GetClientInfoDto() *ClientInfoDto {
	return r._clientInfoDto
}
