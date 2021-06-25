package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加工单-回流单（新接口） APIRequest
alibaba.wdk.ums.handling.get

加工单-回流单（新接口）
*/
type AlibabaWdkUmsHandlingGetRequest struct {
    model.Params

    // 仓库编码
    warehouseCode   string 

}

func NewAlibabaWdkUmsHandlingGetRequest() *AlibabaWdkUmsHandlingGetRequest{
    return &AlibabaWdkUmsHandlingGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsHandlingGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.handling.get"
}

func (r AlibabaWdkUmsHandlingGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsHandlingGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkUmsHandlingGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

