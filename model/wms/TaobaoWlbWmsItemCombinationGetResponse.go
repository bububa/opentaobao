package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询组合商品的组合关系 API返回值 
taobao.wlb.wms.item.combination.get

查询组合商品的组合关系
*/
type TaobaoWlbWmsItemCombinationGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsItemCombinationGetResponse
}

// 查询组合商品的组合关系 成功返回结果
type TaobaoWlbWmsItemCombinationGetResponse struct {
    XMLName xml.Name `xml:"wlb_wms_item_combination_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回结果
    Result   *TaobaoWlbWmsItemCombinationGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
