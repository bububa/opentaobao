package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrandSyncAPIResponse 二手房公司品牌数据同步 API返回值
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
type AlibabaAlihouseExistinghomeBrandSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBrandSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrandSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeBrandSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeBrandSyncAPIResponseModel is 二手房公司品牌数据同步 成功返回结果
type AlibabaAlihouseExistinghomeBrandSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_brand_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBrandSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrandSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeBrandSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrandSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeBrandSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrandSyncAPIResponse
func GetAlibabaAlihouseExistinghomeBrandSyncAPIResponse() *AlibabaAlihouseExistinghomeBrandSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeBrandSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeBrandSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeBrandSyncAPIResponse 将 AlibabaAlihouseExistinghomeBrandSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrandSyncAPIResponse(v *AlibabaAlihouseExistinghomeBrandSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrandSyncAPIResponse.Put(v)
}
