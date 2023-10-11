package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest 医疗器械的码查询信息接口 API请求
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
type AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// NewAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceRequest 初始化AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest对象
func NewAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceRequest() *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest {
	return &AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest) GetCode() string {
	return r._code
}
