package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInvTurnoverQueryAPIResponse 业务库存流水查询 API返回值
// taobao.inv.turnover.query
//
// 业务库存流水
type TaobaoInvTurnoverQueryAPIResponse struct {
	model.CommonResponse
	TaobaoInvTurnoverQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInvTurnoverQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInvTurnoverQueryAPIResponseModel).Reset()
}

// TaobaoInvTurnoverQueryAPIResponseModel is 业务库存流水查询 成功返回结果
type TaobaoInvTurnoverQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inv_turnover_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoInvTurnoverQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInvTurnoverQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoInvTurnoverQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInvTurnoverQueryAPIResponse)
	},
}

// GetTaobaoInvTurnoverQueryAPIResponse 从 sync.Pool 获取 TaobaoInvTurnoverQueryAPIResponse
func GetTaobaoInvTurnoverQueryAPIResponse() *TaobaoInvTurnoverQueryAPIResponse {
	return poolTaobaoInvTurnoverQueryAPIResponse.Get().(*TaobaoInvTurnoverQueryAPIResponse)
}

// ReleaseTaobaoInvTurnoverQueryAPIResponse 将 TaobaoInvTurnoverQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInvTurnoverQueryAPIResponse(v *TaobaoInvTurnoverQueryAPIResponse) {
	v.Reset()
	poolTaobaoInvTurnoverQueryAPIResponse.Put(v)
}
