package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmHistoryOuidGetAPIResponse 根据buyerNick获取ouid API返回值
// taobao.crm.history.ouid.get
//
// 根据buyerNick获取ouid
type TaobaoCrmHistoryOuidGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmHistoryOuidGetAPIResponseModel
}

// TaobaoCrmHistoryOuidGetAPIResponseModel is 根据buyerNick获取ouid 成功返回结果
type TaobaoCrmHistoryOuidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_history_ouid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Data *CrmPrivacyResponse `json:"data,omitempty" xml:"data,omitempty"`
}
