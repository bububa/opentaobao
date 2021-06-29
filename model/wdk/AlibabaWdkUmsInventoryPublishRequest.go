package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
初始化覆盖实物库存 API请求
alibaba.wdk.ums.inventory.publish

先去库存这边查询当前实物库存有多少的量，然后算出来需要增加的量。接下来调用ums原来的入库语义的接口进行库存的增量补充
*/
type AlibabaWdkUmsInventoryPublishRequest struct {
    model.Params
    // 1
    _wdkErpArrivalNotice   *WdkErpArrivalNoticeDto
}

// 初始化AlibabaWdkUmsInventoryPublishRequest对象
func NewAlibabaWdkUmsInventoryPublishRequest() *AlibabaWdkUmsInventoryPublishRequest{
    return &AlibabaWdkUmsInventoryPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.inventory.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WdkErpArrivalNotice Setter
// 1
func (r *AlibabaWdkUmsInventoryPublishRequest) SetWdkErpArrivalNotice(_wdkErpArrivalNotice *WdkErpArrivalNoticeDto) error {
    r._wdkErpArrivalNotice = _wdkErpArrivalNotice
    r.Set("wdk_erp_arrival_notice", _wdkErpArrivalNotice)
    return nil
}

// WdkErpArrivalNotice Getter
func (r AlibabaWdkUmsInventoryPublishRequest) GetWdkErpArrivalNotice() *WdkErpArrivalNoticeDto {
    return r._wdkErpArrivalNotice
}
