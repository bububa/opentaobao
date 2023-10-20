package topoaid

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoCrmHistoryOmidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmHistoryOmidGetAPIResponseModel).Reset()
}

// TaobaoCrmHistoryOmidGetAPIResponseModel is 根据buyerNick获取omid 成功返回结果
type TaobaoCrmHistoryOmidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_history_omid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Data *CrmPrivacyResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmHistoryOmidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoCrmHistoryOmidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmHistoryOmidGetAPIResponse)
	},
}

// GetTaobaoCrmHistoryOmidGetAPIResponse 从 sync.Pool 获取 TaobaoCrmHistoryOmidGetAPIResponse
func GetTaobaoCrmHistoryOmidGetAPIResponse() *TaobaoCrmHistoryOmidGetAPIResponse {
	return poolTaobaoCrmHistoryOmidGetAPIResponse.Get().(*TaobaoCrmHistoryOmidGetAPIResponse)
}

// ReleaseTaobaoCrmHistoryOmidGetAPIResponse 将 TaobaoCrmHistoryOmidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmHistoryOmidGetAPIResponse(v *TaobaoCrmHistoryOmidGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmHistoryOmidGetAPIResponse.Put(v)
}
