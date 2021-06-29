package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount账号信息查询 API请求
taobao.open.account.list

OpenAccount账号信息查询
*/
type TaobaoOpenAccountListRequest struct {
    model.Params
    // Open Account的id列表, 每次最多查询 20 个帐户
    _openAccountIds   []int64
    // ISV自己账号的id列表，isvAccountId和openAccountId二选一必填, 每次最多查询 20 个帐户
    _isvAccountIds   []string
}

// 初始化TaobaoOpenAccountListRequest对象
func NewTaobaoOpenAccountListRequest() *TaobaoOpenAccountListRequest{
    return &TaobaoOpenAccountListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountListRequest) GetApiMethodName() string {
    return "taobao.open.account.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenAccountIds Setter
// Open Account的id列表, 每次最多查询 20 个帐户
func (r *TaobaoOpenAccountListRequest) SetOpenAccountIds(_openAccountIds []int64) error {
    r._openAccountIds = _openAccountIds
    r.Set("open_account_ids", _openAccountIds)
    return nil
}

// OpenAccountIds Getter
func (r TaobaoOpenAccountListRequest) GetOpenAccountIds() []int64 {
    return r._openAccountIds
}
// IsvAccountIds Setter
// ISV自己账号的id列表，isvAccountId和openAccountId二选一必填, 每次最多查询 20 个帐户
func (r *TaobaoOpenAccountListRequest) SetIsvAccountIds(_isvAccountIds []string) error {
    r._isvAccountIds = _isvAccountIds
    r.Set("isv_account_ids", _isvAccountIds)
    return nil
}

// IsvAccountIds Getter
func (r TaobaoOpenAccountListRequest) GetIsvAccountIds() []string {
    return r._isvAccountIds
}
