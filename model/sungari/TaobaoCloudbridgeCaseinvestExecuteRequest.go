package sungari

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
红盾云桥案件协查服务 APIRequest
taobao.cloudbridge.caseinvest.execute

通过API接口直接提供政府部门录入及查询函件服务
*/
type TaobaoCloudbridgeCaseinvestExecuteRequest struct {
    model.Params

    // 方法名称
    apiName   string 

    // 方法参数
    data   string 

}

func NewTaobaoCloudbridgeCaseinvestExecuteRequest() *TaobaoCloudbridgeCaseinvestExecuteRequest{
    return &TaobaoCloudbridgeCaseinvestExecuteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetApiMethodName() string {
    return "taobao.cloudbridge.caseinvest.execute"
}

func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCloudbridgeCaseinvestExecuteRequest) SetApiName(apiName string) error {
    r.apiName = apiName
    r.Set("api_name", apiName)
    return nil
}

func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetApiName() string {
    return r.apiName
}

func (r *TaobaoCloudbridgeCaseinvestExecuteRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetData() string {
    return r.data
}

