package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝open security uid 获取接口 API请求
taobao.opensecurity.uid.get

根据明文 taobao user id 换取 app的 open_uid
*/
type TaobaoOpensecurityUidGetRequest struct {
    model.Params
    // 淘宝用户ID
    _tbUserId   int64
}

// 初始化TaobaoOpensecurityUidGetRequest对象
func NewTaobaoOpensecurityUidGetRequest() *TaobaoOpensecurityUidGetRequest{
    return &TaobaoOpensecurityUidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpensecurityUidGetRequest) GetApiMethodName() string {
    return "taobao.opensecurity.uid.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpensecurityUidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbUserId Setter
// 淘宝用户ID
func (r *TaobaoOpensecurityUidGetRequest) SetTbUserId(_tbUserId int64) error {
    r._tbUserId = _tbUserId
    r.Set("tb_user_id", _tbUserId)
    return nil
}

// TbUserId Getter
func (r TaobaoOpensecurityUidGetRequest) GetTbUserId() int64 {
    return r._tbUserId
}
