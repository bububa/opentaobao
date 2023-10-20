package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest 根据ctei查询状态 API请求
// alibaba.ailabs.tmallgenie.auth.device.status.getbyctei
//
// 提供给电信查询设备在线状态值
type AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest struct {
	model.Params
	// ctei
	_ctei string
}

// NewAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiRequest 初始化AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiRequest() *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.status.getbyctei"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCtei is Ctei Setter
// ctei
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest) SetCtei(_ctei string) error {
	r._ctei = _ctei
	r.Set("ctei", _ctei)
	return nil
}

// GetCtei Ctei Getter
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIRequest) GetCtei() string {
	return r._ctei
}
