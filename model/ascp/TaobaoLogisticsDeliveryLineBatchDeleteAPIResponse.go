package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsdeliverylinebatchdeleteAPIResponse 线路能力删除 API返回值
// taobao.logistics.delivery.line.batch.delete
//
// 线路能力删除
type TaobaologisticsdeliverylinebatchdeleteAPIResponse struct {
	model.CommonResponse
	TaobaologisticsdeliverylinebatchdeleteAPIResponseModel
}

// TaobaologisticsdeliverylinebatchdeleteAPIResponseModel is 线路能力删除 成功返回结果
type TaobaologisticsdeliverylinebatchdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_delivery_line_batch_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 线路能力删除出参
	DeliveryLineBatchDeleteResponse *DeliveryLineBatchDeleteResponse `json:"delivery_line_batch_delete_response,omitempty" xml:"delivery_line_batch_delete_response,omitempty"`
}
