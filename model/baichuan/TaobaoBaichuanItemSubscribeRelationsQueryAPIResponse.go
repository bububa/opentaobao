package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按条件查询订阅关系 API返回值 
taobao.baichuan.item.subscribe.relations.query

按条件查询订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationsQueryAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemSubscribeRelationsQueryAPIResponseModel
}

// 按条件查询订阅关系 成功返回结果
type TaobaoBaichuanItemSubscribeRelationsQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"baichuan_item_subscribe_relations_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeRelationsQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
