package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjPresaleSettlementAddlistAPIResponse 预售结算数据回传 API返回值
// alibaba.mj.presale.settlement.addlist
//
// 用于预售活动结算数据的回传。
type AlibabaMjPresaleSettlementAddlistAPIResponse struct {
	model.CommonResponse
	AlibabaMjPresaleSettlementAddlistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjPresaleSettlementAddlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjPresaleSettlementAddlistAPIResponseModel).Reset()
}

// AlibabaMjPresaleSettlementAddlistAPIResponseModel is 预售结算数据回传 成功返回结果
type AlibabaMjPresaleSettlementAddlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_presale_settlement_addlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjPresaleSettlementAddlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaMjPresaleSettlementAddlistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjPresaleSettlementAddlistAPIResponse)
	},
}

// GetAlibabaMjPresaleSettlementAddlistAPIResponse 从 sync.Pool 获取 AlibabaMjPresaleSettlementAddlistAPIResponse
func GetAlibabaMjPresaleSettlementAddlistAPIResponse() *AlibabaMjPresaleSettlementAddlistAPIResponse {
	return poolAlibabaMjPresaleSettlementAddlistAPIResponse.Get().(*AlibabaMjPresaleSettlementAddlistAPIResponse)
}

// ReleaseAlibabaMjPresaleSettlementAddlistAPIResponse 将 AlibabaMjPresaleSettlementAddlistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjPresaleSettlementAddlistAPIResponse(v *AlibabaMjPresaleSettlementAddlistAPIResponse) {
	v.Reset()
	poolAlibabaMjPresaleSettlementAddlistAPIResponse.Put(v)
}
