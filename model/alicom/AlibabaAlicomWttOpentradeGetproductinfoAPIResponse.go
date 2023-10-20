package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomWttOpentradeGetproductinfoAPIResponse 查询存送产品信息 API返回值
// alibaba.alicom.wtt.opentrade.getproductinfo
//
// 话费宝查询产品信息相关配置
type AlibabaAlicomWttOpentradeGetproductinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomWttOpentradeGetproductinfoAPIResponseModel
}

// AlibabaAlicomWttOpentradeGetproductinfoAPIResponseModel is 查询存送产品信息 成功返回结果
type AlibabaAlicomWttOpentradeGetproductinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_wtt_opentrade_getproductinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
