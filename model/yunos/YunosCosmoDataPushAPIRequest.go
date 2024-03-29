package yunos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosCosmoDataPushAPIRequest COSMO-PUSH模式数据接入 API请求
// yunos.cosmo.data.push
//
// YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
type YunosCosmoDataPushAPIRequest struct {
	model.Params
	// 业务方数据源唯一标识，由COSMO平台颁发
	_appId string
	// 业务方推送数据，List结构的JSON序列化字符串
	_jsonModel string
}

// NewYunosCosmoDataPushRequest 初始化YunosCosmoDataPushAPIRequest对象
func NewYunosCosmoDataPushRequest() *YunosCosmoDataPushAPIRequest {
	return &YunosCosmoDataPushAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosCosmoDataPushAPIRequest) Reset() {
	r._appId = ""
	r._jsonModel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosCosmoDataPushAPIRequest) GetApiMethodName() string {
	return "yunos.cosmo.data.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosCosmoDataPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosCosmoDataPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 业务方数据源唯一标识，由COSMO平台颁发
func (r *YunosCosmoDataPushAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r YunosCosmoDataPushAPIRequest) GetAppId() string {
	return r._appId
}

// SetJsonModel is JsonModel Setter
// 业务方推送数据，List结构的JSON序列化字符串
func (r *YunosCosmoDataPushAPIRequest) SetJsonModel(_jsonModel string) error {
	r._jsonModel = _jsonModel
	r.Set("json_model", _jsonModel)
	return nil
}

// GetJsonModel JsonModel Getter
func (r YunosCosmoDataPushAPIRequest) GetJsonModel() string {
	return r._jsonModel
}

var poolYunosCosmoDataPushAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosCosmoDataPushRequest()
	},
}

// GetYunosCosmoDataPushRequest 从 sync.Pool 获取 YunosCosmoDataPushAPIRequest
func GetYunosCosmoDataPushAPIRequest() *YunosCosmoDataPushAPIRequest {
	return poolYunosCosmoDataPushAPIRequest.Get().(*YunosCosmoDataPushAPIRequest)
}

// ReleaseYunosCosmoDataPushAPIRequest 将 YunosCosmoDataPushAPIRequest 放入 sync.Pool
func ReleaseYunosCosmoDataPushAPIRequest(v *YunosCosmoDataPushAPIRequest) {
	v.Reset()
	poolYunosCosmoDataPushAPIRequest.Put(v)
}
