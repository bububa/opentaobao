package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse 登陆标识信息同步 API返回值
// alibaba.alihouse.existinghome.identifying.sync
//
// 登陆标识信息同步
type AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponseModel is 登陆标识信息同步 成功返回结果
type AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_identifying_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回信息
	Result *AlibabaAlihouseExistinghomeIdentifyingSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse
func GetAlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse() *AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse 将 AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse(v *AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse.Put(v)
}
