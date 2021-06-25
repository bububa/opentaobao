package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓封箱回告 APIRequest
alibaba.wdk.fulfill.warehouse.work.order.sealbox

仓封箱回告箱与包裹的关系
*/
type AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest struct {
    model.Params

    // 同城交付物箱
    sameTownBox   *SameTownBox 

}

func NewAlibabaWdkFulfillWarehouseWorkOrderSealboxRequest() *AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest{
    return &AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.warehouse.work.order.sealbox"
}

func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) SetSameTownBox(sameTownBox *SameTownBox) error {
    r.sameTownBox = sameTownBox
    r.Set("same_town_box", sameTownBox)
    return nil
}

func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) GetSameTownBox() *SameTownBox {
    return r.sameTownBox
}

