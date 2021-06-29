package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询rpId API返回值 
taobao.xhotel.rate.relationshipwithroom.get

某个卖家根据rpId查询所有的gid，可分页，不填分页信息则默认显示第一页。
*/
type TaobaoXhotelRateRelationshipwithroomGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateRelationshipwithroomGetResponse
}

// 查询rpId 成功返回结果
type TaobaoXhotelRateRelationshipwithroomGetResponse struct {
    XMLName xml.Name `xml:"xhotel_rate_relationshipwithroom_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    Gids   []string `json:"gids,omitempty" xml:"gids>string,omitempty"`
    // 根据条件所查询的所有结果的总数量
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
