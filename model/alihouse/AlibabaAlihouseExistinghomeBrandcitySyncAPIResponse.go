package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse 二手房城市品牌店数据同步 API返回值
// alibaba.alihouse.existinghome.brandcity.sync
//
// 二手房城市品牌店数据同步
type AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBrandcitySyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeBrandcitySyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeBrandcitySyncAPIResponseModel is 二手房城市品牌店数据同步 成功返回结果
type AlibabaAlihouseExistinghomeBrandcitySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_brandcity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBrandcitySyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrandcitySyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeBrandcitySyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeBrandcitySyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse
func GetAlibabaAlihouseExistinghomeBrandcitySyncAPIResponse() *AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeBrandcitySyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeBrandcitySyncAPIResponse 将 AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrandcitySyncAPIResponse(v *AlibabaAlihouseExistinghomeBrandcitySyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrandcitySyncAPIResponse.Put(v)
}
