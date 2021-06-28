package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销退单上传 APIResponse
taobao.rdc.aligenius.warehouse.reverse.uploading

主要用于商家上传仓库销退单信息
*/
type TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusWarehouseReverseUploadingResponse
}

type TaobaoRdcAligeniusWarehouseReverseUploadingResponse struct {
    XMLName xml.Name `xml:"rdc_aligenius_warehouse_reverse_uploading_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误描述
    
    FailInfo   string `json:"fail_info,omitempty" xml:"fail_info,omitempty"`

    
    // 错误码
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`

    
    // 是否成功
    
    SuccessFlag   bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`

    
}
