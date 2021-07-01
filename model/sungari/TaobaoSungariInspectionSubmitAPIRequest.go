package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSungariInspectionSubmitAPIRequest
抽检指令录入 API请求
taobao.sungari.inspection.submit

抽检指令录入 */
type TaobaoSungariInspectionSubmitAPIRequest struct {
	model.Params
	// 抽检入参
	_data string
	// 操作类型
	_methodName string
}

// NewTaobaoSungariInspectionSubmitRequest 初始化TaobaoSungariInspectionSubmitAPIRequest对象
func NewTaobaoSungariInspectionSubmitRequest() *TaobaoSungariInspectionSubmitAPIRequest {
	return &TaobaoSungariInspectionSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSungariInspectionSubmitAPIRequest) GetApiMethodName() string {
	return "taobao.sungari.inspection.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSungariInspectionSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Data Setter
// 抽检入参
func (r *TaobaoSungariInspectionSubmitAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// Get Data Getter
func (r TaobaoSungariInspectionSubmitAPIRequest) GetData() string {
	return r._data
}

// Set is MethodName Setter
// 操作类型
func (r *TaobaoSungariInspectionSubmitAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// Get MethodName Getter
func (r TaobaoSungariInspectionSubmitAPIRequest) GetMethodName() string {
	return r._methodName
}
