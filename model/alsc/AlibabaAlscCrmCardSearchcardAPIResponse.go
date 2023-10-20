package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardSearchcardAPIResponse 搜索卡实例列表(支持号段查询) API返回值
// alibaba.alsc.crm.card.searchcard
//
// 搜索卡实例列表(支持号段查询)
type AlibabaAlscCrmCardSearchcardAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardSearchcardAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardSearchcardAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardSearchcardAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardSearchcardAPIResponseModel is 搜索卡实例列表(支持号段查询) 成功返回结果
type AlibabaAlscCrmCardSearchcardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_searchcard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardSearchcardAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardSearchcardAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardSearchcardAPIResponse)
	},
}

// GetAlibabaAlscCrmCardSearchcardAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardSearchcardAPIResponse
func GetAlibabaAlscCrmCardSearchcardAPIResponse() *AlibabaAlscCrmCardSearchcardAPIResponse {
	return poolAlibabaAlscCrmCardSearchcardAPIResponse.Get().(*AlibabaAlscCrmCardSearchcardAPIResponse)
}

// ReleaseAlibabaAlscCrmCardSearchcardAPIResponse 将 AlibabaAlscCrmCardSearchcardAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardSearchcardAPIResponse(v *AlibabaAlscCrmCardSearchcardAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardSearchcardAPIResponse.Put(v)
}
