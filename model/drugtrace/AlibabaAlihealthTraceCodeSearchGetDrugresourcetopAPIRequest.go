package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest 根据码获取码信息 API请求
// alibaba.alihealth.trace.code.search.get.drugresourcetop
//
// 根据码获取码信息
type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest struct {
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

// NewAlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest 初始化AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest对象
func NewAlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest() *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest {
	return &AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) Reset() {
	r._code = ""
	r._token = ""
	r._queryAppName = ""
	r._clientId = ""
	r._deviceUtdid = ""
	r._tbUserId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.trace.code.search.get.drugresourcetop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetCode() string {
	return r._code
}

// SetToken is Token Setter
// 校验值
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetToken() string {
	return r._token
}

// SetQueryAppName is QueryAppName Setter
// 查询app名称
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) SetQueryAppName(_queryAppName string) error {
	r._queryAppName = _queryAppName
	r.Set("query_app_name", _queryAppName)
	return nil
}

// GetQueryAppName QueryAppName Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetQueryAppName() string {
	return r._queryAppName
}

// SetClientId is ClientId Setter
// 客户端ip
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetClientId() string {
	return r._clientId
}

// SetDeviceUtdid is DeviceUtdid Setter
// 设备号
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) SetDeviceUtdid(_deviceUtdid string) error {
	r._deviceUtdid = _deviceUtdid
	r.Set("device_utdid", _deviceUtdid)
	return nil
}

// GetDeviceUtdid DeviceUtdid Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetDeviceUtdid() string {
	return r._deviceUtdid
}

// SetTbUserId is TbUserId Setter
// 用户id
func (r *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) SetTbUserId(_tbUserId int64) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// GetTbUserId TbUserId Getter
func (r AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) GetTbUserId() int64 {
	return r._tbUserId
}

var poolAlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest()
	},
}

// GetAlibabaAlihealthTraceCodeSearchGetDrugresourcetopRequest 从 sync.Pool 获取 AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest
func GetAlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest() *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest {
	return poolAlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest.Get().(*AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest)
}

// ReleaseAlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest 将 AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest(v *AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest.Put(v)
}
