package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店推荐菜信息 APIResponse
taobao.koubei.mall.common.store.display.goods.list

提供查询口碑商圈内的门店推荐菜信息
*/
type TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonStoreDisplayGoodsListResponse
}

type TaobaoKoubeiMallCommonStoreDisplayGoodsListResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_store_display_goods_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // API接口返回的result模型
    
    Result   *TaobaoKoubeiMallCommonStoreDisplayGoodsListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
