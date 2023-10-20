package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse 二手房虚拟店铺数据同步 API返回值
// alibaba.alihouse.existinghome.virtualshop.sync
//
// 二手房虚拟店铺数据同步
type AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponseModel is 二手房虚拟店铺数据同步 成功返回结果
type AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_virtualshop_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeVirtualshopSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse
func GetAlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse() *AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse 将 AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse(v *AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse.Put(v)
}
