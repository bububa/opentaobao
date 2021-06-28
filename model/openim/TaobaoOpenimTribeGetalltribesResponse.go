package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户群列表 APIResponse
taobao.openim.tribe.getalltribes

OPENIM群服务获取用户群列表
*/
type TaobaoOpenimTribeGetalltribesAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribeGetalltribesResponse `json:"openim_tribe_getalltribes_response,omitempty"` 
    TaobaoOpenimTribeGetalltribesResponse
}

/* model for simplify = false
type TaobaoOpenimTribeGetalltribesResponse struct {

    // 群列表信息
    
    TribeInfoList  struct {
        TribeInfo  []TribeInfo `json:"tribe_info,omitempty"`
    } `json:"tribe_info_list,omitempty"`
    

}
*/

type TaobaoOpenimTribeGetalltribesResponse struct {

    // 群列表信息
    TribeInfoList   []TribeInfo `json:"tribe_info_list,omitempty"`

}
