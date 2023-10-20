package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilladdressreachablequeryAPIRequest 地址可达查询 API请求
// cainiao.waybill.address.reachable.query
//
// 地址可达查询
type CainiaowaybilladdressreachablequeryAPIRequest struct {
	model.Params
	// 入参
	_reachableRecommendRequestDto *ReachableRecommendRequestDto
	// 调用方对象
	_clientInfoDto *ClientInfoDto
}

// NewCainiaowaybilladdressreachablequeryRequest 初始化CainiaowaybilladdressreachablequeryAPIRequest对象
func NewCainiaowaybilladdressreachablequeryRequest() *CainiaowaybilladdressreachablequeryAPIRequest {
	return &CainiaowaybilladdressreachablequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilladdressreachablequeryAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.address.reachable.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilladdressreachablequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilladdressreachablequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReachableRecommendRequestDto is ReachableRecommendRequestDto Setter
// 入参
func (r *CainiaowaybilladdressreachablequeryAPIRequest) SetReachableRecommendRequestDto(_reachableRecommendRequestDto *ReachableRecommendRequestDto) error {
	r._reachableRecommendRequestDto = _reachableRecommendRequestDto
	r.Set("reachable_recommend_request_dto", _reachableRecommendRequestDto)
	return nil
}

// GetReachableRecommendRequestDto ReachableRecommendRequestDto Getter
func (r CainiaowaybilladdressreachablequeryAPIRequest) GetReachableRecommendRequestDto() *ReachableRecommendRequestDto {
	return r._reachableRecommendRequestDto
}

// SetClientInfoDto is ClientInfoDto Setter
// 调用方对象
func (r *CainiaowaybilladdressreachablequeryAPIRequest) SetClientInfoDto(_clientInfoDto *ClientInfoDto) error {
	r._clientInfoDto = _clientInfoDto
	r.Set("client_info_dto", _clientInfoDto)
	return nil
}

// GetClientInfoDto ClientInfoDto Getter
func (r CainiaowaybilladdressreachablequeryAPIRequest) GetClientInfoDto() *ClientInfoDto {
	return r._clientInfoDto
}
