package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmExchangeCrowdinstanceDeleteAPIResponse 删除人群实例中的指定买家 API返回值
// taobao.crm.exchange.crowdinstance.delete
//
// 删除人群实例中的指定买家
type TaobaoCrmExchangeCrowdinstanceDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoCrmExchangeCrowdinstanceDeleteAPIResponseModel
}

// TaobaoCrmExchangeCrowdinstanceDeleteAPIResponseModel is 删除人群实例中的指定买家 成功返回结果
type TaobaoCrmExchangeCrowdinstanceDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_exchange_crowdinstance_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作成功
	SubSuccess bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`
}
