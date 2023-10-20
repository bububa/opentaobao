package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresscollectresourcetmsdeleteAPIResponse 上门取退可揽范围删除 API返回值
// taobao.logistics.express.collect.resource.tms.delete
//
// 上门取退可揽范围删除
type TaobaologisticsexpresscollectresourcetmsdeleteAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpresscollectresourcetmsdeleteAPIResponseModel
}

// TaobaologisticsexpresscollectresourcetmsdeleteAPIResponseModel is 上门取退可揽范围删除 成功返回结果
type TaobaologisticsexpresscollectresourcetmsdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_collect_resource_tms_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	CollectResourceDeleteResponse *CollectResourceDeleteResponse `json:"collect_resource_delete_response,omitempty" xml:"collect_resource_delete_response,omitempty"`
}
