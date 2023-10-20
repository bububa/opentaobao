package mei

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmeicrmmembergetbypaycodeAPIResponse 支付码获取会员信息 API返回值
// tmall.mei.crm.member.getbypaycode
//
// 通过支付码获取会员信息
type TmallmeicrmmembergetbypaycodeAPIResponse struct {
	model.CommonResponse
	TmallmeicrmmembergetbypaycodeAPIResponseModel
}

// TmallmeicrmmembergetbypaycodeAPIResponseModel is 支付码获取会员信息 成功返回结果
type TmallmeicrmmembergetbypaycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mei_crm_member_getbypaycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
