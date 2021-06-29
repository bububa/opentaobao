package sungari

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽检指令录入 APIRequest
taobao.sungari.inspection.submit

抽检指令录入
*/
type TaobaoSungariInspectionSubmitRequest struct {
    model.Params

    // 抽检入参
    data   string 

    // 操作类型
    methodName   string 

}

func NewTaobaoSungariInspectionSubmitRequest() *TaobaoSungariInspectionSubmitRequest{
    return &TaobaoSungariInspectionSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSungariInspectionSubmitRequest) GetApiMethodName() string {
    return "taobao.sungari.inspection.submit"
}

func (r TaobaoSungariInspectionSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSungariInspectionSubmitRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r TaobaoSungariInspectionSubmitRequest) GetData() string {
    return r.data
}

func (r *TaobaoSungariInspectionSubmitRequest) SetMethodName(methodName string) error {
    r.methodName = methodName
    r.Set("method_name", methodName)
    return nil
}

func (r TaobaoSungariInspectionSubmitRequest) GetMethodName() string {
    return r.methodName
}

