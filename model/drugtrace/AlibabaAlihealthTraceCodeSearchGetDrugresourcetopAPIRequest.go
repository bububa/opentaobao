package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest 根据码获取码信息 API请求
// alibaba.alihealth.trace.code.search.get.drugresourcetop
//
// 根据码获取码信息
type AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 校验值
	_token string
	// 查询app名称
	_queryAppName string
	// 客户端ip
	_clientId string
	// 设备号
	_deviceUtdid string
	// 用户id
	_tbUserId int64
}

// NewAlibabaalihealthtracecodesearchgetdrugresourcetopRequest 初始化AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest对象
func NewAlibabaalihealthtracecodesearchgetdrugresourcetopRequest() *AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest {
	return &AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.trace.code.search.get.drugresourcetop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetCode() string {
	return r._code
}

// SetToken is Token Setter
// 校验值
func (r *AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetToken() string {
	return r._token
}

// SetQueryAppName is QueryAppName Setter
// 查询app名称
func (r *AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) SetQueryAppName(_queryAppName string) error {
	r._queryAppName = _queryAppName
	r.Set("query_app_name", _queryAppName)
	return nil
}

// GetQueryAppName QueryAppName Getter
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetQueryAppName() string {
	return r._queryAppName
}

// SetClientId is ClientId Setter
// 客户端ip
func (r *AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetClientId() string {
	return r._clientId
}

// SetDeviceUtdid is DeviceUtdid Setter
// 设备号
func (r *AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) SetDeviceUtdid(_deviceUtdid string) error {
	r._deviceUtdid = _deviceUtdid
	r.Set("device_utdid", _deviceUtdid)
	return nil
}

// GetDeviceUtdid DeviceUtdid Getter
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetDeviceUtdid() string {
	return r._deviceUtdid
}

// SetTbUserId is TbUserId Setter
// 用户id
func (r *AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) SetTbUserId(_tbUserId int64) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// GetTbUserId TbUserId Getter
func (r AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest) GetTbUserId() int64 {
	return r._tbUserId
}
