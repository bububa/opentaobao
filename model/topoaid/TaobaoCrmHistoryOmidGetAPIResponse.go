package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmHistoryOmidGetAPIResponse 根据buyerNick获取omid API返回值
// taobao.crm.history.omid.get
//
// 根据buyerNick获取ouid
type TaobaoCrmHistoryOmidGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmHistoryOmidGetAPIResponseModel
}

// TaobaoCrmHistoryOmidGetAPIResponseModel is 根据buyerNick获取omid 成功返回结果
type TaobaoCrmHistoryOmidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_history_omid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Data *CrmPrivacyResponse `json:"data,omitempty" xml:"data,omitempty"`
}
