package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoreachablebatchjudgeAPIRequest 是否派送可达判定批量查询接口 API请求
// cainiao.reachable.batchjudge
//
// 提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型
type CainiaoreachablebatchjudgeAPIRequest struct {
	model.Params
	// 1:快递 2:快运
	_addressType int64
	// 收发信息
	_data *RoutingReachableBatchRequestDto
	// 调用方对象
	_clientInfo *ClientInfoDto
}

// NewCainiaoreachablebatchjudgeRequest 初始化CainiaoreachablebatchjudgeAPIRequest对象
func NewCainiaoreachablebatchjudgeRequest() *CainiaoreachablebatchjudgeAPIRequest {
	return &CainiaoreachablebatchjudgeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoreachablebatchjudgeAPIRequest) GetApiMethodName() string {
	return "cainiao.reachable.batchjudge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoreachablebatchjudgeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoreachablebatchjudgeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddressType is AddressType Setter
// 1:快递 2:快运
func (r *CainiaoreachablebatchjudgeAPIRequest) SetAddressType(_addressType int64) error {
	r._addressType = _addressType
	r.Set("address_type", _addressType)
	return nil
}

// GetAddressType AddressType Getter
func (r CainiaoreachablebatchjudgeAPIRequest) GetAddressType() int64 {
	return r._addressType
}

// SetData is Data Setter
// 收发信息
func (r *CainiaoreachablebatchjudgeAPIRequest) SetData(_data *RoutingReachableBatchRequestDto) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r CainiaoreachablebatchjudgeAPIRequest) GetData() *RoutingReachableBatchRequestDto {
	return r._data
}

// SetClientInfo is ClientInfo Setter
// 调用方对象
func (r *CainiaoreachablebatchjudgeAPIRequest) SetClientInfo(_clientInfo *ClientInfoDto) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r CainiaoreachablebatchjudgeAPIRequest) GetClientInfo() *ClientInfoDto {
	return r._clientInfo
}
