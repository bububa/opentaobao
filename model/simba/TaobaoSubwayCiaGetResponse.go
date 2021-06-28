package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询单元智能出价信息 APIResponse
taobao.subway.cia.get

查询单元智能出价信息
*/
type TaobaoSubwayCiaGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSubwayCiaGetResponse `json:"subway_cia_get_response,omitempty"` 
    TaobaoSubwayCiaGetResponse
}

/* model for simplify = false
type TaobaoSubwayCiaGetResponse struct {

    // 单元智能出价信息
    
    Result  *struct {
        CiaConfig  *CiaConfig `json:"cia_config,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoSubwayCiaGetResponse struct {

    // 单元智能出价信息
    Result   *CiaConfig `json:"result,omitempty"`

}
