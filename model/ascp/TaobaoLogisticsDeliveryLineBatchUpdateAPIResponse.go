package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsdeliverylinebatchupdateAPIResponse 线路能力创建/更新 API返回值
// taobao.logistics.delivery.line.batch.update
//
// 线路能力创建/更新
type TaobaologisticsdeliverylinebatchupdateAPIResponse struct {
	model.CommonResponse
	TaobaologisticsdeliverylinebatchupdateAPIResponseModel
}

// TaobaologisticsdeliverylinebatchupdateAPIResponseModel is 线路能力创建/更新 成功返回结果
type TaobaologisticsdeliverylinebatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_delivery_line_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 线路能力创建/更新出参
	DeliveryLineBatchUpdateResponse *DeliveryLineBatchUpdateResponse `json:"delivery_line_batch_update_response,omitempty" xml:"delivery_line_batch_update_response,omitempty"`
}
