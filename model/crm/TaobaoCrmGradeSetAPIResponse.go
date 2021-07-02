package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGradeSetAPIResponse 卖家设置等级规则 API返回值
// taobao.crm.grade.set
//
// 设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。
type TaobaoCrmGradeSetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGradeSetAPIResponseModel
}

// TaobaoCrmGradeSetAPIResponseModel is 卖家设置等级规则 成功返回结果
type TaobaoCrmGradeSetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grade_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true：成功 false：失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
