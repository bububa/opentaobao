package topoaid

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoCrmHistoryOuidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmHistoryOuidGetAPIResponseModel).Reset()
}

// TaobaoCrmHistoryOuidGetAPIResponseModel is 根据buyerNick获取ouid 成功返回结果
type TaobaoCrmHistoryOuidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_history_ouid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Data *CrmPrivacyResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmHistoryOuidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoCrmHistoryOuidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmHistoryOuidGetAPIResponse)
	},
}

// GetTaobaoCrmHistoryOuidGetAPIResponse 从 sync.Pool 获取 TaobaoCrmHistoryOuidGetAPIResponse
func GetTaobaoCrmHistoryOuidGetAPIResponse() *TaobaoCrmHistoryOuidGetAPIResponse {
	return poolTaobaoCrmHistoryOuidGetAPIResponse.Get().(*TaobaoCrmHistoryOuidGetAPIResponse)
}

// ReleaseTaobaoCrmHistoryOuidGetAPIResponse 将 TaobaoCrmHistoryOuidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmHistoryOuidGetAPIResponse(v *TaobaoCrmHistoryOuidGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmHistoryOuidGetAPIResponse.Put(v)
}
