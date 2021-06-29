package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠详情列表 APIResponse
taobao.promotionmisc.common.item.detail.list.get

查询通用单品优惠详情列表。
*/
type TaobaoPromotionmiscCommonItemDetailListGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscCommonItemDetailListGetResponse
}

type TaobaoPromotionmiscCommonItemDetailListGetResponse struct {
    XMLName xml.Name `xml:"promotionmisc_common_item_detail_list_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否查询成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 数据总数量
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // 活动详情列表
    
    DetailList   []CommonItemDetail `json:"detail_list,omitempty" xml:"detail_list>common_item_detail,omitempty"`
    
    
}
