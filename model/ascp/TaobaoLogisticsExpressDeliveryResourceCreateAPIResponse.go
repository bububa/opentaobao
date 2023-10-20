package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressdeliveryresourcecreateAPIResponse 新建/更新配资源 API返回值
// taobao.logistics.express.delivery.resource.create
//
// 新建/更新配资源
type TaobaologisticsexpressdeliveryresourcecreateAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpressdeliveryresourcecreateAPIResponseModel
}

// TaobaologisticsexpressdeliveryresourcecreateAPIResponseModel is 新建/更新配资源 成功返回结果
type TaobaologisticsexpressdeliveryresourcecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_delivery_resource_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	DeliveryResourceCreateResponse *DeliveryResourceCreateResponse `json:"delivery_resource_create_response,omitempty" xml:"delivery_resource_create_response,omitempty"`
}
