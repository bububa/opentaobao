package waybill

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillAddressReachableQueryAPIRequest) Reset() {
	r._reachableRecommendRequestDto = nil
	r._clientInfoDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillAddressReachableQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.address.reachable.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillAddressReachableQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillAddressReachableQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolCainiaoWaybillAddressReachableQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillAddressReachableQueryRequest()
	},
}

// GetCainiaoWaybillAddressReachableQueryRequest 从 sync.Pool 获取 CainiaoWaybillAddressReachableQueryAPIRequest
func GetCainiaoWaybillAddressReachableQueryAPIRequest() *CainiaoWaybillAddressReachableQueryAPIRequest {
	return poolCainiaoWaybillAddressReachableQueryAPIRequest.Get().(*CainiaoWaybillAddressReachableQueryAPIRequest)
}

// ReleaseCainiaoWaybillAddressReachableQueryAPIRequest 将 CainiaoWaybillAddressReachableQueryAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillAddressReachableQueryAPIRequest(v *CainiaoWaybillAddressReachableQueryAPIRequest) {
	v.Reset()
	poolCainiaoWaybillAddressReachableQueryAPIRequest.Put(v)
}
