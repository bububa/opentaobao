package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商圈内的超值特惠商品信息 APIResponse
taobao.koubei.mall.common.item.super.discount.list

查询商圈超值特惠商品信息列表
*/
type TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonItemSuperDiscountListResponse
}

type TaobaoKoubeiMallCommonItemSuperDiscountListResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_item_super_discount_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoKoubeiMallCommonItemSuperDiscountListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
