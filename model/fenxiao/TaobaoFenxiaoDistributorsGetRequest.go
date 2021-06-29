package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分销商信息 API请求
taobao.fenxiao.distributors.get

查询和当前登录供应商有合作关系的分销商的信息
*/
type TaobaoFenxiaoDistributorsGetRequest struct {
    model.Params
    // 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
    _nicks   string
}

// 初始化TaobaoFenxiaoDistributorsGetRequest对象
func NewTaobaoFenxiaoDistributorsGetRequest() *TaobaoFenxiaoDistributorsGetRequest{
    return &TaobaoFenxiaoDistributorsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDistributorsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.distributors.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDistributorsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nicks Setter
// 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
func (r *TaobaoFenxiaoDistributorsGetRequest) SetNicks(_nicks string) error {
    r._nicks = _nicks
    r.Set("nicks", _nicks)
    return nil
}

// Nicks Getter
func (r TaobaoFenxiaoDistributorsGetRequest) GetNicks() string {
    return r._nicks
}
