package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取open security uid for isv APIResponse
taobao.opensecurity.isv.uid.get

根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联
*/
type TaobaoOpensecurityIsvUidGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpensecurityIsvUidGetResponse `json:"opensecurity_isv_uid_get_response,omitempty"` 
    TaobaoOpensecurityIsvUidGetResponse
}

/* model for simplify = false
type TaobaoOpensecurityIsvUidGetResponse struct {

    // open security tbUserId for ISV，淘宝账号对ISV级别的唯一open security ID，用于同一个ISV多个Appkey间数据共享。
    
    OpenUidIsv   string `json:"open_uid_isv,omitempty"`
    

}
*/

type TaobaoOpensecurityIsvUidGetResponse struct {

    // open security tbUserId for ISV，淘宝账号对ISV级别的唯一open security ID，用于同一个ISV多个Appkey间数据共享。
    OpenUidIsv   string `json:"open_uid_isv,omitempty"`

}
