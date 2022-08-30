package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponse 销售活动解绑预存金商品 API返回值
// alibaba.alihouse.newhome.activity.deposit.unbind
//
// 销售活动解绑预存金商品
type AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponseModel
}

// AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponseModel is 销售活动解绑预存金商品 成功返回结果
type AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_deposit_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文属性
	Result *AlibabaAlihouseNewhomeActivityDepositUnbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
