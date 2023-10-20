package yunos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunoscosmodatapushAPIRequest COSMO-PUSH模式数据接入 API请求
// yunos.cosmo.data.push
//
// YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
type YunoscosmodatapushAPIRequest struct {
	model.Params
	// 业务方数据源唯一标识，由COSMO平台颁发
	_appId string
	// 业务方推送数据，List结构的JSON序列化字符串
	_jsonModel string
}

// NewYunoscosmodatapushRequest 初始化YunoscosmodatapushAPIRequest对象
func NewYunoscosmodatapushRequest() *YunoscosmodatapushAPIRequest {
	return &YunoscosmodatapushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunoscosmodatapushAPIRequest) GetApiMethodName() string {
	return "yunos.cosmo.data.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunoscosmodatapushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunoscosmodatapushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 业务方数据源唯一标识，由COSMO平台颁发
func (r *YunoscosmodatapushAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r YunoscosmodatapushAPIRequest) GetAppId() string {
	return r._appId
}

// SetJsonModel is JsonModel Setter
// 业务方推送数据，List结构的JSON序列化字符串
func (r *YunoscosmodatapushAPIRequest) SetJsonModel(_jsonModel string) error {
	r._jsonModel = _jsonModel
	r.Set("json_model", _jsonModel)
	return nil
}

// GetJsonModel JsonModel Getter
func (r YunoscosmodatapushAPIRequest) GetJsonModel() string {
	return r._jsonModel
}
