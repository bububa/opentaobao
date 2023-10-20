package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest 医疗器械的码查询信息接口 API请求
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
type AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// NewAlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceRequest 初始化AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest对象
func NewAlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceRequest() *AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest {
	return &AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest) GetCode() string {
	return r._code
}
