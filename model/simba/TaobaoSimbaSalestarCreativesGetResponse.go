package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）批量获取创意 APIResponse
taobao.simba.salestar.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
type TaobaoSimbaSalestarCreativesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSalestarCreativesGetResponse `json:"simba_salestar_creatives_get_response,omitempty"` 
    TaobaoSimbaSalestarCreativesGetResponse
}

/* model for simplify = false
type TaobaoSimbaSalestarCreativesGetResponse struct {

    // 创意对象列表
    
    Creatives  struct {
        Creative  []Creative `json:"creative,omitempty"`
    } `json:"creatives,omitempty"`
    

}
*/

type TaobaoSimbaSalestarCreativesGetResponse struct {

    // 创意对象列表
    Creatives   []Creative `json:"creatives,omitempty"`

}
