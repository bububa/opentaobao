package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcBigposBanksaleQueryAPIResponse 大pos银行卡查账接口 API返回值
// alibaba.mj.oc.bigpos.banksale.query
//
// 大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
type AlibabaMjOcBigposBanksaleQueryAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcBigposBanksaleQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcBigposBanksaleQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcBigposBanksaleQueryAPIResponseModel).Reset()
}

// AlibabaMjOcBigposBanksaleQueryAPIResponseModel is 大pos银行卡查账接口 成功返回结果
type AlibabaMjOcBigposBanksaleQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_bigpos_banksale_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据主体
	Datas []AlibabaMjOcBigposBanksaleQueryData `json:"datas,omitempty" xml:"datas>alibaba_mj_oc_bigpos_banksale_query_data,omitempty"`
	// 明细数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcBigposBanksaleQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
	m.Total = 0
}

var poolAlibabaMjOcBigposBanksaleQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcBigposBanksaleQueryAPIResponse)
	},
}

// GetAlibabaMjOcBigposBanksaleQueryAPIResponse 从 sync.Pool 获取 AlibabaMjOcBigposBanksaleQueryAPIResponse
func GetAlibabaMjOcBigposBanksaleQueryAPIResponse() *AlibabaMjOcBigposBanksaleQueryAPIResponse {
	return poolAlibabaMjOcBigposBanksaleQueryAPIResponse.Get().(*AlibabaMjOcBigposBanksaleQueryAPIResponse)
}

// ReleaseAlibabaMjOcBigposBanksaleQueryAPIResponse 将 AlibabaMjOcBigposBanksaleQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcBigposBanksaleQueryAPIResponse(v *AlibabaMjOcBigposBanksaleQueryAPIResponse) {
	v.Reset()
	poolAlibabaMjOcBigposBanksaleQueryAPIResponse.Put(v)
}
