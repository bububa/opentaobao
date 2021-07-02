package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageaccountMesssageMassSendAPIResponse 消息号开放-消息群发 API返回值
// taobao.messageaccount.messsage.mass.send
//
// 外部 isv 调用该进口来进行消息号消息的群发
type TaobaoMessageaccountMesssageMassSendAPIResponse struct {
	model.CommonResponse
	TaobaoMessageaccountMesssageMassSendAPIResponseModel
}

// TaobaoMessageaccountMesssageMassSendAPIResponseModel is 消息号开放-消息群发 成功返回结果
type TaobaoMessageaccountMesssageMassSendAPIResponseModel struct {
	XMLName xml.Name `xml:"messageaccount_messsage_mass_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoMessageaccountMesssageMassSendResult `json:"result,omitempty" xml:"result,omitempty"`
}
