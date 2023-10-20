package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeRestpayfeeModifyAPIResponse 分阶段订单尾款改价 API返回值
// tmall.aliauto.trade.restpayfee.modify
//
// 汽车商家通过此api修改整车分阶段订单的尾款金额
type TmallAliautoTradeRestpayfeeModifyAPIResponse struct {
	model.CommonResponse
	TmallAliautoTradeRestpayfeeModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoTradeRestpayfeeModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoTradeRestpayfeeModifyAPIResponseModel).Reset()
}

// TmallAliautoTradeRestpayfeeModifyAPIResponseModel is 分阶段订单尾款改价 成功返回结果
type TmallAliautoTradeRestpayfeeModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_trade_restpayfee_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoTradeRestpayfeeModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoTradeRestpayfeeModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoTradeRestpayfeeModifyAPIResponse)
	},
}

// GetTmallAliautoTradeRestpayfeeModifyAPIResponse 从 sync.Pool 获取 TmallAliautoTradeRestpayfeeModifyAPIResponse
func GetTmallAliautoTradeRestpayfeeModifyAPIResponse() *TmallAliautoTradeRestpayfeeModifyAPIResponse {
	return poolTmallAliautoTradeRestpayfeeModifyAPIResponse.Get().(*TmallAliautoTradeRestpayfeeModifyAPIResponse)
}

// ReleaseTmallAliautoTradeRestpayfeeModifyAPIResponse 将 TmallAliautoTradeRestpayfeeModifyAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoTradeRestpayfeeModifyAPIResponse(v *TmallAliautoTradeRestpayfeeModifyAPIResponse) {
	v.Reset()
	poolTmallAliautoTradeRestpayfeeModifyAPIResponse.Put(v)
}
