package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotCardInfoAPIResponse
物联卡信息查询 API返回值
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询 */
type AlibabaAliqinFcIotCardInfoAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotCardInfoAPIResponseModel
}

// AlibabaAliqinFcIotCardInfoAPIResponseModel is 物联卡信息查询 成功返回结果
type AlibabaAliqinFcIotCardInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_cardInfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlibabaAliqinFcIotCardInfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
