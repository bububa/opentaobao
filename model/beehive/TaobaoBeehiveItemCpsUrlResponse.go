package beehive

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分佣链接生成接口 APIResponse
taobao.beehive.item.cps.url

传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接
*/
type TaobaoBeehiveItemCpsUrlAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBeehiveItemCpsUrlResponse `json:"beehive_item_cps_url_response,omitempty"` 
    TaobaoBeehiveItemCpsUrlResponse
}

/* model for simplify = false
type TaobaoBeehiveItemCpsUrlResponse struct {

    // 结果对象
    
    Result  *struct {
        TaobaoBeehiveItemCpsUrlResultDo  *TaobaoBeehiveItemCpsUrlResultDo `json:"taobao_beehive_item_cps_url_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBeehiveItemCpsUrlResponse struct {

    // 结果对象
    Result   *TaobaoBeehiveItemCpsUrlResultDo `json:"result,omitempty"`

}
