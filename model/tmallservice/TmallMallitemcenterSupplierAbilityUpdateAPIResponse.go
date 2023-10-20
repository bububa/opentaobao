package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmallitemcentersupplierabilityupdateAPIResponse 门店服务能力授权接口 API返回值
// tmall.mallitemcenter.supplier.ability.update
//
// 门店服务能力授权
type TmallmallitemcentersupplierabilityupdateAPIResponse struct {
	model.CommonResponse
	TmallmallitemcentersupplierabilityupdateAPIResponseModel
}

// TmallmallitemcentersupplierabilityupdateAPIResponseModel is 门店服务能力授权接口 成功返回结果
type TmallmallitemcentersupplierabilityupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mallitemcenter_supplier_ability_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallmallitemcentersupplierabilityupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
