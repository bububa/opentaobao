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
    Response *TaobaoSimbaSerchcrowdGetResponse `json:"taobao_simba_serchcrowd_get_response,omitempty"`
}

type TaobaoSimbaSerchcrowdGetResponse struct {

    // result
    Adgrouptargetingtags   []TaobaoSimbaSerchcrowdGetResult `json:"adgrouptargetingtags,omitempty"`

}
