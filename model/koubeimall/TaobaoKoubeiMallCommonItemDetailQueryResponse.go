package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品详情信息 APIResponse
taobao.koubei.mall.common.item.detail.query

查询口碑综合体内商品详情信息
*/
type TaobaoKoubeiMallCommonItemDetailQueryAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonItemDetailQueryResponse
}

type TaobaoKoubeiMallCommonItemDetailQueryResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_item_detail_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoKoubeiMallCommonItemDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
