package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
初始化覆盖实物库存 APIRequest
alibaba.wdk.ums.inventory.publish

先去库存这边查询当前实物库存有多少的量，然后算出来需要增加的量。接下来调用ums原来的入库语义的接口进行库存的增量补充
*/
type AlibabaWdkUmsInventoryPublishRequest struct {
    model.Params

    // 1
    wdkErpArrivalNotice   *WdkErpArrivalNoticeDto 

}

func NewAlibabaWdkUmsInventoryPublishRequest() *AlibabaWdkUmsInventoryPublishRequest{
    return &AlibabaWdkUmsInventoryPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsInventoryPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inventory.publish"
}

func (r AlibabaWdkUmsInventoryPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsInventoryPublishRequest) SetWdkErpArrivalNotice(wdkErpArrivalNotice *WdkErpArrivalNoticeDto) error {
    r.wdkErpArrivalNotice = wdkErpArrivalNotice
    r.Set("wdk_erp_arrival_notice", wdkErpArrivalNotice)
    return nil
}

func (r AlibabaWdkUmsInventoryPublishRequest) GetWdkErpArrivalNotice() *WdkErpArrivalNoticeDto {
    return r.wdkErpArrivalNotice
}

