package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotRechargeCardAPIResponse 按终端号订购增值业务 API返回值
// alibaba.aliqin.fc.iot.rechargeCard
//
// 按终端号订购增值业务
type AlibabaAliqinFcIotRechargeCardAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotRechargeCardAPIResponseModel
}

// AlibabaAliqinFcIotRechargeCardAPIResponseModel is 按终端号订购增值业务 成功返回结果
type AlibabaAliqinFcIotRechargeCardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_rechargeCard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotRechargeCardResult `json:"result,omitempty" xml:"result,omitempty"`
}
