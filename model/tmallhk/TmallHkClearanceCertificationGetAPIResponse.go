package tmallhk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearancecertificationgetAPIResponse 获取订单清关材料实名信息 API返回值
// tmall.hk.clearance.certification.get
//
// 获取订单清关材料实名信息
type TmallhkclearancecertificationgetAPIResponse struct {
	model.CommonResponse
	TmallhkclearancecertificationgetAPIResponseModel
}

// TmallhkclearancecertificationgetAPIResponseModel is 获取订单清关材料实名信息 成功返回结果
type TmallhkclearancecertificationgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_hk_clearance_certification_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示信息
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 错误编码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 返回内容数据
	Data *OrderCertifyResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 接口是否返回成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
