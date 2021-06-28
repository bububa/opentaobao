package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝open security uid 获取接口 APIResponse
taobao.opensecurity.uid.get

根据明文 taobao user id 换取 app的 open_uid
*/
type TaobaoOpensecurityUidGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpensecurityUidGetResponse `json:"opensecurity_uid_get_response,omitempty"` 
    TaobaoOpensecurityUidGetResponse
}

/* model for simplify = false
type TaobaoOpensecurityUidGetResponse struct {

    // open security tbUserId，淘宝用户对每个Appkey会有唯一的一个open_uid
    
    OpenUid   string `json:"open_uid,omitempty"`
    

}
*/

type TaobaoOpensecurityUidGetResponse struct {

    // open security tbUserId，淘宝用户对每个Appkey会有唯一的一个open_uid
    OpenUid   string `json:"open_uid,omitempty"`

}
