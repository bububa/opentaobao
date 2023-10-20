package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtpaymerchantfundstypemodifyAPIResponse 修改摊位分账类型 API返回值
// tmall.nrt.pay.merchant.fundstype.modify
//
// 修改摊位分账类型
type TmallnrtpaymerchantfundstypemodifyAPIResponse struct {
	model.CommonResponse
	TmallnrtpaymerchantfundstypemodifyAPIResponseModel
}

// TmallnrtpaymerchantfundstypemodifyAPIResponseModel is 修改摊位分账类型 成功返回结果
type TmallnrtpaymerchantfundstypemodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_pay_merchant_fundstype_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统参数
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
