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
type AlibabaAlihealthMedicalbaseOrderStatusSyncRequest struct {
    model.Params
    // 订单信息
    _orderlSyncDTO   *OrderlSyncDTO
}

// 初始化AlibabaAlihealthMedicalbaseOrderStatusSyncRequest对象
func NewAlibabaAlihealthMedicalbaseOrderStatusSyncRequest() *AlibabaAlihealthMedicalbaseOrderStatusSyncRequest{
    return &AlibabaAlihealthMedicalbaseOrderStatusSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medicalbase.order.status.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderlSyncDTO Setter
// 订单信息
func (r *AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) SetOrderlSyncDTO(_orderlSyncDTO *OrderlSyncDTO) error {
    r._orderlSyncDTO = _orderlSyncDTO
    r.Set("orderl_sync_d_t_o", _orderlSyncDTO)
    return nil
}

// OrderlSyncDTO Getter
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) GetOrderlSyncDTO() *OrderlSyncDTO {
    return r._orderlSyncDTO
}
