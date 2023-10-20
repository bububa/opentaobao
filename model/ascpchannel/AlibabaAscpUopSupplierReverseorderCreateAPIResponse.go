package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierreverseordercreateAPIResponse 商家ERP发起创建销退单服务 API返回值
// alibaba.ascp.uop.supplier.reverseorder.create
//
// 商家在收到消费者实物退货后，在ERP发起创建销退单服务
type AlibabaascpuopsupplierreverseordercreateAPIResponse struct {
	model.CommonResponse
	AlibabaascpuopsupplierreverseordercreateAPIResponseModel
}

// AlibabaascpuopsupplierreverseordercreateAPIResponseModel is 商家ERP发起创建销退单服务 成功返回结果
type AlibabaascpuopsupplierreverseordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_reverseorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	ReverseCreateResponse *ResultWrapper `json:"reverse_create_response,omitempty" xml:"reverse_create_response,omitempty"`
}
