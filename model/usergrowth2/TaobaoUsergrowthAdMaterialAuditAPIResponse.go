package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthAdMaterialAuditAPIResponse 素材审核 API返回值
// taobao.usergrowth.ad.material.audit
//
// 素材审核
type TaobaoUsergrowthAdMaterialAuditAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthAdMaterialAuditAPIResponseModel
}

// TaobaoUsergrowthAdMaterialAuditAPIResponseModel is 素材审核 成功返回结果
type TaobaoUsergrowthAdMaterialAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_ad_material_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 参数错误
	ResponseCode int64 `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 新增创意id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
