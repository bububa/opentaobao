package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosungariinspectionsubmitAPIRequest 抽检指令录入 API请求
// taobao.sungari.inspection.submit
//
// 抽检指令录入
type TaobaosungariinspectionsubmitAPIRequest struct {
	model.Params
	// 抽检入参
	_data string
	// 操作类型
	_methodName string
}

// NewTaobaosungariinspectionsubmitRequest 初始化TaobaosungariinspectionsubmitAPIRequest对象
func NewTaobaosungariinspectionsubmitRequest() *TaobaosungariinspectionsubmitAPIRequest {
	return &TaobaosungariinspectionsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosungariinspectionsubmitAPIRequest) GetApiMethodName() string {
	return "taobao.sungari.inspection.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosungariinspectionsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosungariinspectionsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 抽检入参
func (r *TaobaosungariinspectionsubmitAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaosungariinspectionsubmitAPIRequest) GetData() string {
	return r._data
}

// SetMethodName is MethodName Setter
// 操作类型
func (r *TaobaosungariinspectionsubmitAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r TaobaosungariinspectionsubmitAPIRequest) GetMethodName() string {
	return r._methodName
}
