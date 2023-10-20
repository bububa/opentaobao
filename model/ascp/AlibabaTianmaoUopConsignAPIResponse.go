package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaouopconsignAPIResponse 阿里巴巴.天猫. 履约订单. 发货 API返回值
// alibaba.tianmao.uop.consign
//
// 阿里巴巴.天猫. 履约订单. 发货
type AlibabatianmaouopconsignAPIResponse struct {
	model.CommonResponse
	AlibabatianmaouopconsignAPIResponseModel
}

// AlibabatianmaouopconsignAPIResponseModel is 阿里巴巴.天猫. 履约订单. 发货 成功返回结果
type AlibabatianmaouopconsignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_uop_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
