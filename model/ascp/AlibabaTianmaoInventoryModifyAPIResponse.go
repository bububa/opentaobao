package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoInventoryModifyAPIResponse 阿里巴巴.天猫.aic库存.修改 API返回值
// alibaba.tianmao.inventory.modify
//
// 阿里巴巴.天猫.aic库存.修改
type AlibabaTianmaoInventoryModifyAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoInventoryModifyAPIResponseModel
}

// AlibabaTianmaoInventoryModifyAPIResponseModel is 阿里巴巴.天猫.aic库存.修改 成功返回结果
type AlibabaTianmaoInventoryModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_inventory_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
