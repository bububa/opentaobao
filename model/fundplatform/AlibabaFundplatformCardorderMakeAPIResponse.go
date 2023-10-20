package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderMakeAPIResponse 通知制卡商制卡 API返回值
// alibaba.fundplatform.cardorder.make
//
// 该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。
type AlibabaFundplatformCardorderMakeAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformCardorderMakeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardorderMakeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformCardorderMakeAPIResponseModel).Reset()
}

// AlibabaFundplatformCardorderMakeAPIResponseModel is 通知制卡商制卡 成功返回结果
type AlibabaFundplatformCardorderMakeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorder_make_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Response *AlibabaFundplatformCardorderMakeStruct `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardorderMakeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolAlibabaFundplatformCardorderMakeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformCardorderMakeAPIResponse)
	},
}

// GetAlibabaFundplatformCardorderMakeAPIResponse 从 sync.Pool 获取 AlibabaFundplatformCardorderMakeAPIResponse
func GetAlibabaFundplatformCardorderMakeAPIResponse() *AlibabaFundplatformCardorderMakeAPIResponse {
	return poolAlibabaFundplatformCardorderMakeAPIResponse.Get().(*AlibabaFundplatformCardorderMakeAPIResponse)
}

// ReleaseAlibabaFundplatformCardorderMakeAPIResponse 将 AlibabaFundplatformCardorderMakeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformCardorderMakeAPIResponse(v *AlibabaFundplatformCardorderMakeAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformCardorderMakeAPIResponse.Put(v)
}
