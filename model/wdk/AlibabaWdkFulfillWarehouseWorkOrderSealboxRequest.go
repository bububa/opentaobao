package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓封箱回告 API请求
alibaba.wdk.fulfill.warehouse.work.order.sealbox

仓封箱回告箱与包裹的关系
*/
type AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest struct {
    model.Params
    // 同城交付物箱
    sameTownBox   *SameTownBox
}

// 初始化AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest对象
func NewAlibabaWdkFulfillWarehouseWorkOrderSealboxRequest() *AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest{
    return &AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.warehouse.work.order.sealbox"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SameTownBox Setter
// 同城交付物箱
func (r *AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) SetSameTownBox(sameTownBox *SameTownBox) error {
    r.sameTownBox = sameTownBox
    r.Set("same_town_box", sameTownBox)
    return nil
}

// SameTownBox Getter
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxRequest) GetSameTownBox() *SameTownBox {
    return r.sameTownBox
}
