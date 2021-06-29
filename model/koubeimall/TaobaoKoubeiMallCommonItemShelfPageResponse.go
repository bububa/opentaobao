package koubeimall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店货架商品列表信息查询 API返回值 
taobao.koubei.mall.common.item.shelf.page

查询口碑综合体内门店货架商品列表信息接口
*/
type TaobaoKoubeiMallCommonItemShelfPageAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiMallCommonItemShelfPageResponse
}

// 门店货架商品列表信息查询 成功返回结果
type TaobaoKoubeiMallCommonItemShelfPageResponse struct {
    XMLName xml.Name `xml:"koubei_mall_common_item_shelf_page_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // API接口返回的result模型
    Result   *TaobaoKoubeiMallCommonItemShelfPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
