package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoReachableBatchjudgeAPIRequest 是否派送可达判定批量查询接口 API请求
// cainiao.reachable.batchjudge
//
// 提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型
type CainiaoReachableBatchjudgeAPIRequest struct {
	model.Params
	// 调用方对象
	_clientInfo *ClientInfoDto
	// 1:快递 2:快运
	_addressType int64
	// 收发信息
	_data *RoutingReachableBatchRequestDto
}

// NewCainiaoReachableBatchjudgeRequest 初始化CainiaoReachableBatchjudgeAPIRequest对象
func NewCainiaoReachableBatchjudgeRequest() *CainiaoReachableBatchjudgeAPIRequest {
	return &CainiaoReachableBatchjudgeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoReachableBatchjudgeAPIRequest) GetApiMethodName() string {
	return "cainiao.reachable.batchjudge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoReachableBatchjudgeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetClientInfo is ClientInfo Setter
// 调用方对象
func (r *CainiaoReachableBatchjudgeAPIRequest) SetClientInfo(_clientInfo *ClientInfoDto) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r CainiaoReachableBatchjudgeAPIRequest) GetClientInfo() *ClientInfoDto {
	return r._clientInfo
}

// SetAddressType is AddressType Setter
// 1:快递 2:快运
func (r *CainiaoReachableBatchjudgeAPIRequest) SetAddressType(_addressType int64) error {
	r._addressType = _addressType
	r.Set("address_type", _addressType)
	return nil
}

// GetAddressType AddressType Getter
func (r CainiaoReachableBatchjudgeAPIRequest) GetAddressType() int64 {
	return r._addressType
}

// SetData is Data Setter
// 收发信息
func (r *CainiaoReachableBatchjudgeAPIRequest) SetData(_data *RoutingReachableBatchRequestDto) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r CainiaoReachableBatchjudgeAPIRequest) GetData() *RoutingReachableBatchRequestDto {
	return r._data
}
