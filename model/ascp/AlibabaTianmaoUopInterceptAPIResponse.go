package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaouopinterceptAPIResponse 阿里巴巴.天猫. 履约订单. 配送拦截 API返回值
// alibaba.tianmao.uop.intercept
//
// 阿里巴巴.天猫. 履约订单. 配送拦截
type AlibabatianmaouopinterceptAPIResponse struct {
	model.CommonResponse
	AlibabatianmaouopinterceptAPIResponseModel
}

// AlibabatianmaouopinterceptAPIResponseModel is 阿里巴巴.天猫. 履约订单. 配送拦截 成功返回结果
type AlibabatianmaouopinterceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_uop_intercept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
