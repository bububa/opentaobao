package sungari

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
红盾云桥案件协查服务 API请求
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

// 初始化TaobaoCloudbridgeCaseinvestExecuteRequest对象
func NewTaobaoCloudbridgeCaseinvestExecuteRequest() *TaobaoCloudbridgeCaseinvestExecuteRequest{
    return &TaobaoCloudbridgeCaseinvestExecuteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetApiMethodName() string {
    return "taobao.cloudbridge.caseinvest.execute"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApiName Setter
// 方法名称
func (r *TaobaoCloudbridgeCaseinvestExecuteRequest) SetApiName(apiName string) error {
    r.apiName = apiName
    r.Set("api_name", apiName)
    return nil
}

// ApiName Getter
func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetApiName() string {
    return r.apiName
}
// Data Setter
// 方法参数
func (r *TaobaoCloudbridgeCaseinvestExecuteRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r TaobaoCloudbridgeCaseinvestExecuteRequest) GetData() string {
    return r.data
}
