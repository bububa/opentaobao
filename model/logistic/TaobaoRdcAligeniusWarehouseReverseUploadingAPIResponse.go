package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponse 销退单上传 API返回值
// taobao.rdc.aligenius.warehouse.reverse.uploading
//
// 主要用于商家上传仓库销退单信息
type TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponseModel
}

// TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponseModel is 销退单上传 成功返回结果
type TaobaoRdcAligeniusWarehouseReverseUploadingAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_warehouse_reverse_uploading_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	FailInfo string `json:"fail_info,omitempty" xml:"fail_info,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}
