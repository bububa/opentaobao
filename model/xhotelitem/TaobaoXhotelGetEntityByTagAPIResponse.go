package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据标签查询实体 API返回值 
taobao.xhotel.get.entity.by.tag

根据标签查询实体
*/
type TaobaoXhotelGetEntityByTagAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelGetEntityByTagAPIResponseModel
}

// 根据标签查询实体 成功返回结果
type TaobaoXhotelGetEntityByTagAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_get_entity_by_tag_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    TagQueryResult   *TagQueryResult `json:"tag_query_result,omitempty" xml:"tag_query_result,omitempty"`
}
