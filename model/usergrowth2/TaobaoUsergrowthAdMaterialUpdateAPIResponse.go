package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthadmaterialupdateAPIResponse 素材更新 API返回值
// taobao.usergrowth.ad.material.update
//
// 素材更新
type TaobaousergrowthadmaterialupdateAPIResponse struct {
	model.CommonResponse
	TaobaousergrowthadmaterialupdateAPIResponseModel
}

// TaobaousergrowthadmaterialupdateAPIResponseModel is 素材更新 成功返回结果
type TaobaousergrowthadmaterialupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_ad_material_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 更新创意的id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 参数错误
	ResponseCode int64 `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 请求结果
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
