package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsconsignresendAPIResponse 修改物流公司和运单号 API返回值
// alibaba.ascp.logistics.consign.resend
//
// 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：
// 1、必须是已发货订单，自己联系发货的必须50天内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
// 2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司
type AlibabaascplogisticsconsignresendAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticsconsignresendAPIResponseModel
}

// AlibabaascplogisticsconsignresendAPIResponseModel is 修改物流公司和运单号 成功返回结果
type AlibabaascplogisticsconsignresendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_consign_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaascplogisticsconsignresendResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
