package oversea

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOverseaExchagerateGetAPIResponse 汇率信息获取 API返回值
// alibaba.oversea.exchagerate.get
//
// 提供外部汇率查询接口
type AlibabaOverseaExchagerateGetAPIResponse struct {
	model.CommonResponse
	AlibabaOverseaExchagerateGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOverseaExchagerateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOverseaExchagerateGetAPIResponseModel).Reset()
}

// AlibabaOverseaExchagerateGetAPIResponseModel is 汇率信息获取 成功返回结果
type AlibabaOverseaExchagerateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_oversea_exchagerate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果描述
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOverseaExchagerateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOverseaExchagerateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOverseaExchagerateGetAPIResponse)
	},
}

// GetAlibabaOverseaExchagerateGetAPIResponse 从 sync.Pool 获取 AlibabaOverseaExchagerateGetAPIResponse
func GetAlibabaOverseaExchagerateGetAPIResponse() *AlibabaOverseaExchagerateGetAPIResponse {
	return poolAlibabaOverseaExchagerateGetAPIResponse.Get().(*AlibabaOverseaExchagerateGetAPIResponse)
}

// ReleaseAlibabaOverseaExchagerateGetAPIResponse 将 AlibabaOverseaExchagerateGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOverseaExchagerateGetAPIResponse(v *AlibabaOverseaExchagerateGetAPIResponse) {
	v.Reset()
	poolAlibabaOverseaExchagerateGetAPIResponse.Put(v)
}
