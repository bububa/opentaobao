package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentbookticketconfirmvtwoAPIResponse 火车票代理商接口——确认出票是否成功v2--增加鉴权校验 API返回值
// taobao.train.agent.bookticket.confirm.vtwo
//
// 火车票代理商接口——确认出票是否成功
type TaobaotrainagentbookticketconfirmvtwoAPIResponse struct {
	model.CommonResponse
	TaobaotrainagentbookticketconfirmvtwoAPIResponseModel
}

// TaobaotrainagentbookticketconfirmvtwoAPIResponseModel is 火车票代理商接口——确认出票是否成功v2--增加鉴权校验 成功返回结果
type TaobaotrainagentbookticketconfirmvtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_bookticket_confirm_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
