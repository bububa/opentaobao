package miniappcloud

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudStoreListfileAPIResponse 云存储根据文件名反查地址 API返回值
// taobao.miniapp.cloud.store.listfile
//
// 云存储中，根据文件名反查地址
type TaobaoMiniappCloudStoreListfileAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappCloudStoreListfileAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappCloudStoreListfileAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappCloudStoreListfileAPIResponseModel).Reset()
}

// TaobaoMiniappCloudStoreListfileAPIResponseModel is 云存储根据文件名反查地址 成功返回结果
type TaobaoMiniappCloudStoreListfileAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_cloud_store_listfile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Files []File `json:"files,omitempty" xml:"files>file,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappCloudStoreListfileAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Files = m.Files[:0]
}

var poolTaobaoMiniappCloudStoreListfileAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappCloudStoreListfileAPIResponse)
	},
}

// GetTaobaoMiniappCloudStoreListfileAPIResponse 从 sync.Pool 获取 TaobaoMiniappCloudStoreListfileAPIResponse
func GetTaobaoMiniappCloudStoreListfileAPIResponse() *TaobaoMiniappCloudStoreListfileAPIResponse {
	return poolTaobaoMiniappCloudStoreListfileAPIResponse.Get().(*TaobaoMiniappCloudStoreListfileAPIResponse)
}

// ReleaseTaobaoMiniappCloudStoreListfileAPIResponse 将 TaobaoMiniappCloudStoreListfileAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappCloudStoreListfileAPIResponse(v *TaobaoMiniappCloudStoreListfileAPIResponse) {
	v.Reset()
	poolTaobaoMiniappCloudStoreListfileAPIResponse.Put(v)
}
