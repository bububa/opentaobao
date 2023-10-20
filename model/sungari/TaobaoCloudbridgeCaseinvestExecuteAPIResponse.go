package sungari

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocloudbridgecaseinvestexecuteAPIResponse 红盾云桥案件协查服务 API返回值
// taobao.cloudbridge.caseinvest.execute
//
// 通过API接口直接提供政府部门录入及查询函件服务
type TaobaocloudbridgecaseinvestexecuteAPIResponse struct {
	model.CommonResponse
	TaobaocloudbridgecaseinvestexecuteAPIResponseModel
}

// TaobaocloudbridgecaseinvestexecuteAPIResponseModel is 红盾云桥案件协查服务 成功返回结果
type TaobaocloudbridgecaseinvestexecuteAPIResponseModel struct {
	XMLName xml.Name `xml:"cloudbridge_caseinvest_execute_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaocloudbridgecaseinvestexecuteResult `json:"result,omitempty" xml:"result,omitempty"`
}
