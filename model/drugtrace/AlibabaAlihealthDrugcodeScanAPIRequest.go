package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodescanAPIRequest 查询扫码信息 API请求
// alibaba.alihealth.drugcode.scan
//
// 查询扫码信息
type AlibabaalihealthdrugcodescanAPIRequest struct {
	model.Params
	// 20位码
	_code string
	// 渠道
	_queryAppName string
	// 用户ip
	_clientId string
	// 设备标识
	_deviceUtdid string
	// 用户ID
	_userId string
}

// NewAlibabaalihealthdrugcodescanRequest 初始化AlibabaalihealthdrugcodescanAPIRequest对象
func NewAlibabaalihealthdrugcodescanRequest() *AlibabaalihealthdrugcodescanAPIRequest {
	return &AlibabaalihealthdrugcodescanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodescanAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodescanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodescanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 20位码
func (r *AlibabaalihealthdrugcodescanAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugcodescanAPIRequest) GetCode() string {
	return r._code
}

// SetQueryAppName is QueryAppName Setter
// 渠道
func (r *AlibabaalihealthdrugcodescanAPIRequest) SetQueryAppName(_queryAppName string) error {
	r._queryAppName = _queryAppName
	r.Set("query_app_name", _queryAppName)
	return nil
}

// GetQueryAppName QueryAppName Getter
func (r AlibabaalihealthdrugcodescanAPIRequest) GetQueryAppName() string {
	return r._queryAppName
}

// SetClientId is ClientId Setter
// 用户ip
func (r *AlibabaalihealthdrugcodescanAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaalihealthdrugcodescanAPIRequest) GetClientId() string {
	return r._clientId
}

// SetDeviceUtdid is DeviceUtdid Setter
// 设备标识
func (r *AlibabaalihealthdrugcodescanAPIRequest) SetDeviceUtdid(_deviceUtdid string) error {
	r._deviceUtdid = _deviceUtdid
	r.Set("device_utdid", _deviceUtdid)
	return nil
}

// GetDeviceUtdid DeviceUtdid Getter
func (r AlibabaalihealthdrugcodescanAPIRequest) GetDeviceUtdid() string {
	return r._deviceUtdid
}

// SetUserId is UserId Setter
// 用户ID
func (r *AlibabaalihealthdrugcodescanAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaalihealthdrugcodescanAPIRequest) GetUserId() string {
	return r._userId
}
