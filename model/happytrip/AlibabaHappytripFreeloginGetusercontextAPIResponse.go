package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripFreeloginGetusercontextAPIResponse 提供给外部系统的免登校验 API返回值
// alibaba.happytrip.freelogin.getusercontext
//
// 免登融合，提供免登相关接口给外部供应商做登录验证
type AlibabaHappytripFreeloginGetusercontextAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripFreeloginGetusercontextAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripFreeloginGetusercontextAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripFreeloginGetusercontextAPIResponseModel).Reset()
}

// AlibabaHappytripFreeloginGetusercontextAPIResponseModel is 提供给外部系统的免登校验 成功返回结果
type AlibabaHappytripFreeloginGetusercontextAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_freelogin_getusercontext_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求响应
	Rep *AlibabaHappytripFreeloginGetusercontextResult `json:"rep,omitempty" xml:"rep,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripFreeloginGetusercontextAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rep = nil
}

var poolAlibabaHappytripFreeloginGetusercontextAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripFreeloginGetusercontextAPIResponse)
	},
}

// GetAlibabaHappytripFreeloginGetusercontextAPIResponse 从 sync.Pool 获取 AlibabaHappytripFreeloginGetusercontextAPIResponse
func GetAlibabaHappytripFreeloginGetusercontextAPIResponse() *AlibabaHappytripFreeloginGetusercontextAPIResponse {
	return poolAlibabaHappytripFreeloginGetusercontextAPIResponse.Get().(*AlibabaHappytripFreeloginGetusercontextAPIResponse)
}

// ReleaseAlibabaHappytripFreeloginGetusercontextAPIResponse 将 AlibabaHappytripFreeloginGetusercontextAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripFreeloginGetusercontextAPIResponse(v *AlibabaHappytripFreeloginGetusercontextAPIResponse) {
	v.Reset()
	poolAlibabaHappytripFreeloginGetusercontextAPIResponse.Put(v)
}
