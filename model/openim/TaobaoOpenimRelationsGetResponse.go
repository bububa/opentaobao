package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openim账号的聊天关系 APIResponse
taobao.openim.relations.get

获取openim账号的聊天关系
*/
type TaobaoOpenimRelationsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_relations_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户信息列表
    
    Users   []OpenImUser `json:"users,omitempty" xml:"