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
type TaobaoSungariInspectionSubmitRequest struct {
    model.Params
    // 抽检入参
    _data   string
    // 操作类型
    _methodName   string
}

// 初始化TaobaoSungariInspectionSubmitRequest对象
func NewTaobaoSungariInspectionSubmitRequest() *TaobaoSungariInspectionSubmitRequest{
    return &TaobaoSungariInspectionSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSungariInspectionSubmitRequest) GetApiMethodName() string {
    return "taobao.sungari.inspection.submit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSungariInspectionSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Data Setter
// 抽检入参
func (r *TaobaoSungariInspectionSubmitRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r TaobaoSungariInspectionSubmitRequest) GetData() string {
    return r._data
}
// MethodName Setter
// 操作类型
func (r *TaobaoSungariInspectionSubmitRequest) SetMethodName(_methodName string) error {
    r._methodName = _methodName
    r.Set("method_name", _methodName)
    return nil
}

// MethodName Getter
func (r TaobaoSungariInspectionSubmitRequest) GetMethodName() string {
    return r._methodName
}
