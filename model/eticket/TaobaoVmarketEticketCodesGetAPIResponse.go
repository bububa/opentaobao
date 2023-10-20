package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketcodesgetAPIResponse 电子凭证码列表查询 API返回值
// taobao.vmarket.eticket.codes.get
//
// 查询某个订单的所有码的列表
type TaobaovmarketeticketcodesgetAPIResponse struct {
	model.CommonResponse
	TaobaovmarketeticketcodesgetAPIResponseModel
}

// TaobaovmarketeticketcodesgetAPIResponseModel is 电子凭证码列表查询 成功返回结果
type TaobaovmarketeticketcodesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_codes_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 电子凭证码列表
	EticketCodes []EticketCode `json:"eticket_codes,omitempty" xml:"eticket_codes>eticket_code,omitempty"`
	// 记录总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
