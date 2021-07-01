package sungari

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽检指令录入 API请求
taobao.sungari.inspection.submit

抽检指令录入
*/
type TaobaoSungariInspectionSubmitAPIRequest struct {
    model.Params
    // 抽检入参
    _data   string
    // 操作类型
    _methodName   string
}

// 初始化TaobaoSungariInspectionSubmitAPIRequest对象
func NewTaobaoSungariInspectionSubmitRequest() *TaobaoSungariInspectionSubmitAPIRequest{
    return &TaobaoSungariInspectionSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSungariInspectionSubmitAPIRequest) GetApiMethodName() string {
    return "taobao.sungari.inspection.submit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSungariInspectionSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Data Setter
// 抽检入参
func (r *TaobaoSungariInspectionSubmitAPIRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r TaobaoSungariInspectionSubmitAPIRequest) GetData() string {
    return r._data
}
// MethodName Setter
// 操作类型
func (r *TaobaoSungariInspectionSubmitAPIRequest) SetMethodName(_methodName string) error {
    r._methodName = _methodName
    r.Set("method_name", _methodName)
    return nil
}

// MethodName Getter
func (r TaobaoSungariInspectionSubmitAPIRequest) GetMethodName() string {
    return r._methodName
}
