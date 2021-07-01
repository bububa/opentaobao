package miniappcloud

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
MongoDB插入单条数据 API返回值 
taobao.miniapp.cloud.mongo.insert

向商家应用云中插入一条记录，用于外部数据同步到应用中
*/
type TaobaoMiniappCloudMongoInsertAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappCloudMongoInsertAPIResponseModel
}

// MongoDB插入单条数据 成功返回结果
type TaobaoMiniappCloudMongoInsertAPIResponseModel struct {
    XMLName xml.Name `xml:"miniapp_cloud_mongo_insert_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回的数据ID，在表中为_id字段
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
}
