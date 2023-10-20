package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudStoreRelationAddAPIResponse 云存储添加 API返回值
// taobao.miniapp.cloud.store.relation.add
//
// 用于用户上传文件之后回写云存储的关联关系
type TaobaoMiniappCloudStoreRelationAddAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappCloudStoreRelationAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappCloudStoreRelationAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappCloudStoreRelationAddAPIResponseModel).Reset()
}

// TaobaoMiniappCloudStoreRelationAddAPIResponseModel is 云存储添加 成功返回结果
type TaobaoMiniappCloudStoreRelationAddAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_cloud_store_relation_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 文件真实url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 云存储文件唯一ID
	FileId string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappCloudStoreRelationAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Url = ""
	m.FileId = ""
}

var poolTaobaoMiniappCloudStoreRelationAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappCloudStoreRelationAddAPIResponse)
	},
}

// GetTaobaoMiniappCloudStoreRelationAddAPIResponse 从 sync.Pool 获取 TaobaoMiniappCloudStoreRelationAddAPIResponse
func GetTaobaoMiniappCloudStoreRelationAddAPIResponse() *TaobaoMiniappCloudStoreRelationAddAPIResponse {
	return poolTaobaoMiniappCloudStoreRelationAddAPIResponse.Get().(*TaobaoMiniappCloudStoreRelationAddAPIResponse)
}

// ReleaseTaobaoMiniappCloudStoreRelationAddAPIResponse 将 TaobaoMiniappCloudStoreRelationAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappCloudStoreRelationAddAPIResponse(v *TaobaoMiniappCloudStoreRelationAddAPIResponse) {
	v.Reset()
	poolTaobaoMiniappCloudStoreRelationAddAPIResponse.Put(v)
}
