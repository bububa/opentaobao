package miniappcloud

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新MongoDB中的数据 API返回值 
taobao.miniapp.cloud.mongo.update

更新MongoDB中的数据
*/
type TaobaoMiniappCloudMongoUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappCloudMongoUpdateResponse
}

// 更新MongoDB中的数据 成功返回结果
type TaobaoMiniappCloudMongoUpdateResponse struct {
    XMLName xml.Name `xml:"miniapp_cloud_mongo_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 更新的记录数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
