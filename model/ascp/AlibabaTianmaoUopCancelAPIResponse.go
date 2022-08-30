package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopCancelAPIResponse 阿里巴巴.天猫. 履约订单. 取消 API返回值
// alibaba.tianmao.uop.cancel
//
// 阿里巴巴.天猫. 履约订单. 取消
type AlibabaTianmaoUopCancelAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoUopCancelAPIResponseModel
}

// AlibabaTianmaoUopCancelAPIResponseModel is 阿里巴巴.天猫. 履约订单. 取消 成功返回结果
type AlibabaTianmaoUopCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_uop_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
