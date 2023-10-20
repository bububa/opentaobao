package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse 落地页获取 API返回值
// alibaba.alihouse.newhome.link.info.obtain
//
// 落地页获取
type AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel is 落地页获取 成功返回结果
type AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_link_info_obtain_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeLinkInfoObtainResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeLinkInfoObtainAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeLinkInfoObtainAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse
func GetAlibabaAlihouseNewhomeLinkInfoObtainAPIResponse() *AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse {
	return poolAlibabaAlihouseNewhomeLinkInfoObtainAPIResponse.Get().(*AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeLinkInfoObtainAPIResponse 将 AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeLinkInfoObtainAPIResponse(v *AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeLinkInfoObtainAPIResponse.Put(v)
}
