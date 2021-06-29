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
type TaobaoOpenAccountDeleteRequest struct {
    model.Params
    // Open Account的id列表
    openAccountIds   []int64
    // ISV自己账号的id列表
    isvAccountIds   []string
}

// 初始化TaobaoOpenAccountDeleteRequest对象
func NewTaobaoOpenAccountDeleteRequest() *TaobaoOpenAccountDeleteRequest{
    return &TaobaoOpenAccountDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountDeleteRequest) GetApiMethodName() string {
    return "taobao.open.account.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenAccountIds Setter
// Open Account的id列表
func (r *TaobaoOpenAccountDeleteRequest) SetOpenAccountIds(openAccountIds []int64) error {
    r.openAccountIds = openAccountIds
    r.Set("open_account_ids", openAccountIds)
    return nil
}

// OpenAccountIds Getter
func (r TaobaoOpenAccountDeleteRequest) GetOpenAccountIds() []int64 {
    return r.openAccountIds
}
// IsvAccountIds Setter
// ISV自己账号的id列表
func (r *TaobaoOpenAccountDeleteRequest) SetIsvAccountIds(isvAccountIds []string) error {
    r.isvAccountIds = isvAccountIds
    r.Set("isv_account_ids", isvAccountIds)
    return nil
}

// IsvAccountIds Getter
func (r TaobaoOpenAccountDeleteRequest) GetIsvAccountIds() []string {
    return r.isvAccountIds
}
