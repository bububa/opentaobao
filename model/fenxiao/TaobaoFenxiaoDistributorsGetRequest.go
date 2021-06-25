package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分销商信息 APIRequest
taobao.fenxiao.distributors.get

查询和当前登录供应商有合作关系的分销商的信息
*/
type TaobaoFenxiaoDistributorsGetRequest struct {
    model.Params

    // 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
    nicks   string 

}

func NewTaobaoFenxiaoDistributorsGetRequest() *TaobaoFenxiaoDistributorsGetRequest{
    return &TaobaoFenxiaoDistributorsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoDistributorsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.distributors.get"
}

func (r TaobaoFenxiaoDistributorsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoDistributorsGetRequest) SetNicks(nicks string) error {
    r.nicks = nicks
    r.Set("nicks", nicks)
    return nil
}

func (r TaobaoFenxiaoDistributorsGetRequest) GetNicks() string {
    return r.nicks
}

