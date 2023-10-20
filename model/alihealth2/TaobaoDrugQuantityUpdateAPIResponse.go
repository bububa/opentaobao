package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugQuantityUpdateAPIResponse 商家更新库存 API返回值
// taobao.drug.quantity.update
//
// 商家通过top接口可以直接修改商品库存
type TaobaoDrugQuantityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoDrugQuantityUpdateAPIResponseModel
}

// TaobaoDrugQuantityUpdateAPIResponseModel is 商家更新库存 成功返回结果
type TaobaoDrugQuantityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_quantity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoDrugQuantityUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
