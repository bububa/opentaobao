package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeScanAPIRequest 查询扫码信息 API请求
// alibaba.alihealth.drugcode.scan
//
// 查询扫码信息
type AlibabaAlihealthDrugcodeScanAPIRequest struct {
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

// NewAlibabaAlihealthDrugcodeScanRequest 初始化AlibabaAlihealthDrugcodeScanAPIRequest对象
func NewAlibabaAlihealthDrugcodeScanRequest() *AlibabaAlihealthDrugcodeScanAPIRequest {
	return &AlibabaAlihealthDrugcodeScanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 20位码
func (r *AlibabaAlihealthDrugcodeScanAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetCode() string {
	return r._code
}

// SetQueryAppName is QueryAppName Setter
// 渠道
func (r *AlibabaAlihealthDrugcodeScanAPIRequest) SetQueryAppName(_queryAppName string) error {
	r._queryAppName = _queryAppName
	r.Set("query_app_name", _queryAppName)
	return nil
}

// GetQueryAppName QueryAppName Getter
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetQueryAppName() string {
	return r._queryAppName
}

// SetClientId is ClientId Setter
// 用户ip
func (r *AlibabaAlihealthDrugcodeScanAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetClientId() string {
	return r._clientId
}

// SetDeviceUtdid is DeviceUtdid Setter
// 设备标识
func (r *AlibabaAlihealthDrugcodeScanAPIRequest) SetDeviceUtdid(_deviceUtdid string) error {
	r._deviceUtdid = _deviceUtdid
	r.Set("device_utdid", _deviceUtdid)
	return nil
}

// GetDeviceUtdid DeviceUtdid Getter
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetDeviceUtdid() string {
	return r._deviceUtdid
}

// SetUserId is UserId Setter
// 用户ID
func (r *AlibabaAlihealthDrugcodeScanAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlihealthDrugcodeScanAPIRequest) GetUserId() string {
	return r._userId
}
