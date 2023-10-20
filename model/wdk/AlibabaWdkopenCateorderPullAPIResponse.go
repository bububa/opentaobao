package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkopenCateorderPullAPIResponse 商户回传餐饮加工单状态 API返回值
// alibaba.wdkopen.cateorder.pull
//
// 商户回传餐饮加工单状态
type AlibabaWdkopenCateorderPullAPIResponse struct {
	model.CommonResponse
	AlibabaWdkopenCateorderPullAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkopenCateorderPullAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkopenCateorderPullAPIResponseModel).Reset()
}

// AlibabaWdkopenCateorderPullAPIResponseModel is 商户回传餐饮加工单状态 成功返回结果
type AlibabaWdkopenCateorderPullAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkopen_cateorder_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用返回
	TopBaseResult *TopBaseResult `json:"top_base_result,omitempty" xml:"top_base_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkopenCateorderPullAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopBaseResult = nil
}

var poolAlibabaWdkopenCateorderPullAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkopenCateorderPullAPIResponse)
	},
}

// GetAlibabaWdkopenCateorderPullAPIResponse 从 sync.Pool 获取 AlibabaWdkopenCateorderPullAPIResponse
func GetAlibabaWdkopenCateorderPullAPIResponse() *AlibabaWdkopenCateorderPullAPIResponse {
	return poolAlibabaWdkopenCateorderPullAPIResponse.Get().(*AlibabaWdkopenCateorderPullAPIResponse)
}

// ReleaseAlibabaWdkopenCateorderPullAPIResponse 将 AlibabaWdkopenCateorderPullAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkopenCateorderPullAPIResponse(v *AlibabaWdkopenCateorderPullAPIResponse) {
	v.Reset()
	poolAlibabaWdkopenCateorderPullAPIResponse.Put(v)
}
