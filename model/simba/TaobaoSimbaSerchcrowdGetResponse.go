package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据推广单元id获取搜索溢价人群 APIResponse
taobao.simba.serchcrowd.get

根据推广单元id获取搜索溢价人群
*/
type TaobaoSimbaSerchcrowdGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSerchcrowdGetResponse `json:"simba_serchcrowd_get_response,omitempty"` 
    TaobaoSimbaSerchcrowdGetResponse
}

/* model for simplify = false
type TaobaoSimbaSerchcrowdGetResponse struct {

    // result
    
    Adgrouptargetingtags  struct {
        TaobaoSimbaSerchcrowdGetResult  []TaobaoSimbaSerchcrowdGetResult `json:"taobao_simba_serchcrowd_get_result,omitempty"`
    } `json:"adgrouptargetingtags,omitempty"`
    

}
*/

type TaobaoSimbaSerchcrowdGetResponse struct {

    // result
    Adgrouptargetingtags   []TaobaoSimbaSerchcrowdGetResult `json:"adgrouptargetingtags,omitempty"`

}
