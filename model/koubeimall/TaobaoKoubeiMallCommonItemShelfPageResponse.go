package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店货架商品列表信息查询 APIResponse
taobao.koubei.mall.common.item.shelf.page

查询口碑综合体内门店货架商品列表信息接口
*/
type TaobaoKoubeiMallCommonItemShelfPageAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonItemShelfPageResponse
}

type TaobaoKoubeiMallCommonItemShelfPageResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_item_shelf_page_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // API接口返回的result模型
    
    Result   *TaobaoKoubeiMallCommonItemShelfPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
