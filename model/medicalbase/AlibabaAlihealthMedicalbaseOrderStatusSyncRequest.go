package medicalbase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
号源直连订单状态同步接口 APIRequest
alibaba.alihealth.medicalbase.order.status.sync

互联网医院isv批量通过接口批量导入
*/
type AlibabaAlihealthMedicalbaseOrderStatusSyncRequest struct {
    model.Params

    // 订单信息
    orderlSyncDTO   *OrderlSyncDto 

}

func NewAlibabaAlihealthMedicalbaseOrderStatusSyncRequest() *AlibabaAlihealthMedicalbaseOrderStatusSyncRequest{
    return &AlibabaAlihealthMedicalbaseOrderStatusSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medicalbase.order.status.sync"
}

func (r AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) SetOrderlSyncDTO(orderlSyncDTO *OrderlSyncDto) error {
    r.orderlSyncDTO = orderlSyncDTO
    r.Set("orderl_sync_d_t_o", orderlSyncDTO)
    return nil
}

func (r AlibabaAlihealthMedicalbaseOrderStatusSyncRequest) GetOrderlSyncDTO() *OrderlSyncDto {
    return r.orderlSyncDTO
}

