package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse 大pos银行卡调账申请 API返回值
// alibaba.mj.oc.bigpos.banksale.adjustment.apply
//
// 大pos银行卡调账申请
type AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponseModel).Reset()
}

// AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponseModel is 大pos银行卡调账申请 成功返回结果
type AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_bigpos_banksale_adjustment_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolAlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse)
	},
}

// GetAlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse 从 sync.Pool 获取 AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse
func GetAlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse() *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse {
	return poolAlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse.Get().(*AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse)
}

// ReleaseAlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse 将 AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse(v *AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse) {
	v.Reset()
	poolAlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse.Put(v)
}
