package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse 查询底导数据 API返回值
// alibaba.alihealth.pregnancy.navigateinfo.query
//
// 备孕管理--获取底部导航信息
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponseModel is 查询底导数据 成功返回结果
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_navigateinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse)
	},
}

// GetAlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse
func GetAlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse() *AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse {
	return poolAlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse.Get().(*AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse)
}

// ReleaseAlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse 将 AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse(v *AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse.Put(v)
}
