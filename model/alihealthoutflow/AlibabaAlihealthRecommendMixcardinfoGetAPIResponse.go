package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRecommendMixcardinfoGetAPIResponse 快应用混合卡片信息 API返回值
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
type AlibabaAlihealthRecommendMixcardinfoGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthRecommendMixcardinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthRecommendMixcardinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthRecommendMixcardinfoGetAPIResponseModel).Reset()
}

// AlibabaAlihealthRecommendMixcardinfoGetAPIResponseModel is 快应用混合卡片信息 成功返回结果
type AlibabaAlihealthRecommendMixcardinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_recommend_mixcardinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	CardResult *ServiceResult `json:"card_result,omitempty" xml:"card_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthRecommendMixcardinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CardResult = nil
}

var poolAlibabaAlihealthRecommendMixcardinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthRecommendMixcardinfoGetAPIResponse)
	},
}

// GetAlibabaAlihealthRecommendMixcardinfoGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthRecommendMixcardinfoGetAPIResponse
func GetAlibabaAlihealthRecommendMixcardinfoGetAPIResponse() *AlibabaAlihealthRecommendMixcardinfoGetAPIResponse {
	return poolAlibabaAlihealthRecommendMixcardinfoGetAPIResponse.Get().(*AlibabaAlihealthRecommendMixcardinfoGetAPIResponse)
}

// ReleaseAlibabaAlihealthRecommendMixcardinfoGetAPIResponse 将 AlibabaAlihealthRecommendMixcardinfoGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthRecommendMixcardinfoGetAPIResponse(v *AlibabaAlihealthRecommendMixcardinfoGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthRecommendMixcardinfoGetAPIResponse.Put(v)
}
