package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse 天猫好房店铺装修-Astore上翻 API返回值
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
type AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel is 天猫好房店铺装修-Astore上翻 成功返回结果
type AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopconfig_astore_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *AlibabaAlihouseNewhomeShopconfigAstoreSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse
func GetAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse() *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse 将 AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse(v *AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse.Put(v)
}
