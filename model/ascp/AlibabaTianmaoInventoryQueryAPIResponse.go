package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoInventoryQueryAPIResponse 阿里巴巴.天猫.aic库存.查询 API返回值
// alibaba.tianmao.inventory.query
//
// 阿里巴巴.天猫.aic库存.查询
type AlibabaTianmaoInventoryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoInventoryQueryAPIResponseModel
}

// AlibabaTianmaoInventoryQueryAPIResponseModel is 阿里巴巴.天猫.aic库存.查询 成功返回结果
type AlibabaTianmaoInventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ListResult *ListResult `json:"list_result,omitempty" xml:"list_result,omitempty"`
}
