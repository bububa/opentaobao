package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询组合商品的组合关系 APIResponse
taobao.wlb.wms.item.combination.get

查询组合商品的组合关系
*/
type TaobaoWlbWmsItemCombinationGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_item_combination_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回结果
    
    Result   *TaobaoWlbWmsItemCombinationGetResult `json:"result,omitempty" xml:"