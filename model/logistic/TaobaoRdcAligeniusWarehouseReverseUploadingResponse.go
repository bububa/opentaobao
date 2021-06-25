package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销退单上传 APIResponse
taobao.rdc.aligenius.warehouse.reverse.uploading

主要用于商家上传仓库销退单信息
*/
type TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRdcAligeniusWarehouseReverseUploadingResponse `json:"taobao_rdc_aligenius_warehouse_reverse_uploading_response,omitempty"`
}

type TaobaoRdcAligeniusWarehouseReverseUploadingResponse struct {

    // 错误描述
    FailInfo   string `json:"fail_info,omitempty"`

    // 错误码
    FailCode   string `json:"fail_code,omitempty"`

    // 是否成功
    SuccessFlag   bool `json:"success_flag,omitempty"`

}
