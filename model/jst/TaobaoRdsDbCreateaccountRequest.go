package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库账户 API请求
taobao.rds.db.createaccount

rds创建数据库账户
*/
type TaobaoRdsDbCreateaccountRequest struct {
    model.Params
    // 入参对象
    _param0   *RequestDbAccountModel
}

// 初始化TaobaoRdsDbCreateaccountRequest对象
func NewTaobaoRdsDbCreateaccountRequest() *TaobaoRdsDbCreateaccountRequest{
    return &TaobaoRdsDbCreateaccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdsDbCreateaccountRequest) GetApiMethodName() string {
    return "taobao.rds.db.createaccount"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdsDbCreateaccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参对象
func (r *TaobaoRdsDbCreateaccountRequest) SetParam0(_param0 *RequestDbAccountModel) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoRdsDbCreateaccountRequest) GetParam0() *RequestDbAccountModel {
    return r._param0
}
