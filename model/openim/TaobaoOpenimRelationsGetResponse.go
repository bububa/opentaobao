package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取openim账号的聊天关系 APIResponse
taobao.openim.relations.get

获取openim账号的聊天关系
*/
type TaobaoOpenimRelationsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimRelationsGetResponse `json:"openim_relations_get_response,omitempty"` 
    TaobaoOpenimRelationsGetResponse
}

/* model for simplify = false
type TaobaoOpenimRelationsGetResponse struct {

    // 用户信息列表
    
    Users  struct {
        OpenImUser  []OpenImUser `json:"open_im_user,omitempty"`
    } `json:"users,omitempty"`
    

}
*/

type TaobaoOpenimRelationsGetResponse struct {

    // 用户信息列表
    Users   []OpenImUser `json:"users,omitempty"`

}
