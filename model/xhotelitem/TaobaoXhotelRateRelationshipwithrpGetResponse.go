package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据gid查询卖家下所有的rpId APIResponse
taobao.xhotel.rate.relationshipwithrp.get

根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据
*/
type TaobaoXhotelRateRelationshipwithrpGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateRelationshipwithrpGetResponse
}

type TaobaoXhotelRateRelationshipwithrpGetResponse struct {
    XMLName xml.Name `xml:"xhotel_rate_relationshipwithrp_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 所查询出的结果，是一个字符串数组
    
    RpIds   []string `json:"rp_ids,omitempty" xml:"rp_ids>string,omitempty"`
    
    
    // 根据条件所查询的所有结果的总数量
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
