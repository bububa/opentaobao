package medicalbase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
号源直连订单状态同步接口 API请求
alibaba.alihealth.medicalbase.order.status.sync

互联网医院isv批量通过接口批量导入
*/
type AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest struct {
    model.Params
    // 订单信息
    _orderlSyncDTO   *OrderlSyncDTO
}

// 初始化AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseOrderStatusSyncRequest() *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest{
    return &AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medicalbase.order.status.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderlSyncDTO Setter
// 订单信息
func (r *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) SetOrderlSyncDTO(_orderlSyncDTO *OrderlSyncDTO) error {
    r._orderlSyncDTO = _orderlSyncDTO
    r.Set("orderl_sync_d_t_o", _orderlSyncDTO)
    return nil
}

// OrderlSyncDTO Getter
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetOrderlSyncDTO() *OrderlSyncDTO {
    return r._orderlSyncDTO
}
