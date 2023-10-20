package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivityDepositBindAPIResponse 销售活动绑定预存金商品id API返回值
// alibaba.alihouse.newhome.activity.deposit.bind
//
// 销售活动绑定预存金商品id
type AlibabaAlihouseNewhomeActivityDepositBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeActivityDepositBindAPIResponseModel
}

// AlibabaAlihouseNewhomeActivityDepositBindAPIResponseModel is 销售活动绑定预存金商品id 成功返回结果
type AlibabaAlihouseNewhomeActivityDepositBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_deposit_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文属性
	Result *AlibabaAlihouseNewhomeActivityDepositBindResult `json:"result,omitempty" xml:"result,omitempty"`
}
