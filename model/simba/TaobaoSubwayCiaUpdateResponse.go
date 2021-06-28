package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量修改单元智能出价 APIResponse
taobao.subway.cia.update

批量修改直通车推广单元的智能出价配置
*/
type TaobaoSubwayCiaUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSubwayCiaUpdateResponse `json:"subway_cia_update_response,omitempty"` 
    TaobaoSubwayCiaUpdateResponse
}

/* model for simplify = false
type TaobaoSubwayCiaUpdateResponse struct {

    // 推广组Id列表
    
    AdgroupList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"adgroup_list,omitempty"`
    

}
*/

type TaobaoSubwayCiaUpdateResponse struct {

    // 推广组Id列表
    AdgroupList   []int64 `json:"adgroup_list,omitempty"`

}
