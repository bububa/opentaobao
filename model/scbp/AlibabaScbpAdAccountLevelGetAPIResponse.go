package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdAccountLevelGetAPIResponse 查询推广账户等级 API返回值
// alibaba.scbp.ad.account.level.get
//
// 查询推广账户等级
type AlibabaScbpAdAccountLevelGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdAccountLevelGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdAccountLevelGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdAccountLevelGetAPIResponseModel).Reset()
}

// AlibabaScbpAdAccountLevelGetAPIResponseModel is 查询推广账户等级 成功返回结果
type AlibabaScbpAdAccountLevelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_account_level_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广账户等级
	CustLevelDto *TopCustLevelDto `json:"cust_level_dto,omitempty" xml:"cust_level_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdAccountLevelGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CustLevelDto = nil
}

var poolAlibabaScbpAdAccountLevelGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdAccountLevelGetAPIResponse)
	},
}

// GetAlibabaScbpAdAccountLevelGetAPIResponse 从 sync.Pool 获取 AlibabaScbpAdAccountLevelGetAPIResponse
func GetAlibabaScbpAdAccountLevelGetAPIResponse() *AlibabaScbpAdAccountLevelGetAPIResponse {
	return poolAlibabaScbpAdAccountLevelGetAPIResponse.Get().(*AlibabaScbpAdAccountLevelGetAPIResponse)
}

// ReleaseAlibabaScbpAdAccountLevelGetAPIResponse 将 AlibabaScbpAdAccountLevelGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdAccountLevelGetAPIResponse(v *AlibabaScbpAdAccountLevelGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdAccountLevelGetAPIResponse.Put(v)
}
