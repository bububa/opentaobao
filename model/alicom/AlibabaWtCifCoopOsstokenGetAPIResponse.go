package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWtCifCoopOsstokenGetAPIResponse 获取oss签名接口 API返回值
// alibaba.wt.cif.coop.osstoken.get
//
// 商家合作上传oss图片获取token接口
type AlibabaWtCifCoopOsstokenGetAPIResponse struct {
	model.CommonResponse
	AlibabaWtCifCoopOsstokenGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWtCifCoopOsstokenGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWtCifCoopOsstokenGetAPIResponseModel).Reset()
}

// AlibabaWtCifCoopOsstokenGetAPIResponseModel is 获取oss签名接口 成功返回结果
type AlibabaWtCifCoopOsstokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wt_cif_coop_osstoken_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWtCifCoopOsstokenGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWtCifCoopOsstokenGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWtCifCoopOsstokenGetAPIResponse)
	},
}

// GetAlibabaWtCifCoopOsstokenGetAPIResponse 从 sync.Pool 获取 AlibabaWtCifCoopOsstokenGetAPIResponse
func GetAlibabaWtCifCoopOsstokenGetAPIResponse() *AlibabaWtCifCoopOsstokenGetAPIResponse {
	return poolAlibabaWtCifCoopOsstokenGetAPIResponse.Get().(*AlibabaWtCifCoopOsstokenGetAPIResponse)
}

// ReleaseAlibabaWtCifCoopOsstokenGetAPIResponse 将 AlibabaWtCifCoopOsstokenGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWtCifCoopOsstokenGetAPIResponse(v *AlibabaWtCifCoopOsstokenGetAPIResponse) {
	v.Reset()
	poolAlibabaWtCifCoopOsstokenGetAPIResponse.Put(v)
}
