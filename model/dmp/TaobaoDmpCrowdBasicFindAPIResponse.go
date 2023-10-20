package dmp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaodmpcrowdbasicfindAPIResponse DMP_BP版人群列表查询 API返回值
// taobao.dmp.crowd.basic.find
//
// DMP_BP版人群列表查询
type TaobaodmpcrowdbasicfindAPIResponse struct {
	model.CommonResponse
	TaobaodmpcrowdbasicfindAPIResponseModel
}

// TaobaodmpcrowdbasicfindAPIResponseModel is DMP_BP版人群列表查询 成功返回结果
type TaobaodmpcrowdbasicfindAPIResponseModel struct {
	XMLName xml.Name `xml:"dmp_crowd_basic_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 人群对象
	Result []CrowdDto `json:"result,omitempty" xml:"result>crowd_dto,omitempty"`
	// 错误码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分页对象
	Pager *Pager `json:"pager,omitempty" xml:"pager,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
