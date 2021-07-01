package miniapp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云存储添加 API返回值 
taobao.miniapp.cloud.store.relation.add

用于用户上传文件之后回写云存储的关联关系
*/
type TaobaoMiniappCloudStoreRelationAddAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappCloudStoreRelationAddAPIResponseModel
}

// 云存储添加 成功返回结果
type TaobaoMiniappCloudStoreRelationAddAPIResponseModel struct {
    XMLName xml.Name `xml:"miniapp_cloud_store_relation_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 文件真实url
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 云存储文件唯一ID
    FileId   string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}
