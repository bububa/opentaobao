package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount删除数据 API请求
taobao.open.account.delete

OpenAccount删除数据
*/
type TaobaoOpenAccountDeleteAPIRequest struct {
    model.Params
    // Open Account的id列表
    _openAccountIds   []int64
    // ISV自己账号的id列表
    _isvAccountIds   []string
}

// 初始化TaobaoOpenAccountDeleteAPIRequest对象
func NewTaobaoOpenAccountDeleteRequest() *TaobaoOpenAccountDeleteAPIRequest{
    return &TaobaoOpenAccountDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.open.account.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenAccountIds Setter
// Open Account的id列表
func (r *TaobaoOpenAccountDeleteAPIRequest) SetOpenAccountIds(_openAccountIds []int64) error {
    r._openAccountIds = _openAccountIds
    r.Set("open_account_ids", _openAccountIds)
    return nil
}

// OpenAccountIds Getter
func (r TaobaoOpenAccountDeleteAPIRequest) GetOpenAccountIds() []int64 {
    return r._openAccountIds
}
// IsvAccountIds Setter
// ISV自己账号的id列表
func (r *TaobaoOpenAccountDeleteAPIRequest) SetIsvAccountIds(_isvAccountIds []string) error {
    r._isvAccountIds = _isvAccountIds
    r.Set("isv_account_ids", _isvAccountIds)
    return nil
}

// IsvAccountIds Getter
func (r TaobaoOpenAccountDeleteAPIRequest) GetIsvAccountIds() []string {
    return r._isvAccountIds
}
