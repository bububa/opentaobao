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
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_common_item_detail_list_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否查询成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"