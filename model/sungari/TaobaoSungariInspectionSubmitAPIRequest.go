package sungari

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSungariInspectionSubmitAPIRequest 抽检指令录入 API请求
// taobao.sungari.inspection.submit
//
// 抽检指令录入
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSungariInspectionSubmitAPIRequest) Reset() {
	r._data = ""
	r._methodName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSungariInspectionSubmitAPIRequest) GetApiMethodName() string {
	return "taobao.sungari.inspection.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSungariInspectionSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSungariInspectionSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 抽检入参
func (r *TaobaoSungariInspectionSubmitAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaoSungariInspectionSubmitAPIRequest) GetData() string {
	return r._data
}

// SetMethodName is MethodName Setter
// 操作类型
func (r *TaobaoSungariInspectionSubmitAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r TaobaoSungariInspectionSubmitAPIRequest) GetMethodName() string {
	return r._methodName
}

var poolTaobaoSungariInspectionSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSungariInspectionSubmitRequest()
	},
}

// GetTaobaoSungariInspectionSubmitRequest 从 sync.Pool 获取 TaobaoSungariInspectionSubmitAPIRequest
func GetTaobaoSungariInspectionSubmitAPIRequest() *TaobaoSungariInspectionSubmitAPIRequest {
	return poolTaobaoSungariInspectionSubmitAPIRequest.Get().(*TaobaoSungariInspectionSubmitAPIRequest)
}

// ReleaseTaobaoSungariInspectionSubmitAPIRequest 将 TaobaoSungariInspectionSubmitAPIRequest 放入 sync.Pool
func ReleaseTaobaoSungariInspectionSubmitAPIRequest(v *TaobaoSungariInspectionSubmitAPIRequest) {
	v.Reset()
	poolTaobaoSungariInspectionSubmitAPIRequest.Put(v)
}
