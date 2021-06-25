package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过hscode获取计量单位 APIRequest
tmall.item.hscode.detail.get

通过hscode获取计量单位和销售单位
*/
type TmallItemHscodeDetailGetRequest struct {
    model.Params

    // hscode
    hscode   string 

}

func NewTmallItemHscodeDetailGetRequest() *TmallItemHscodeDetailGetRequest{
    return &TmallItemHscodeDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemHscodeDetailGetRequest) GetApiMethodName() string {
    return "tmall.item.hscode.detail.get"
}

func (r TmallItemHscodeDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemHscodeDetailGetRequest) SetHscode(hscode string) error {
    r.hscode = hscode
    r.Set("hscode", hscode)
    return nil
}

func (r TmallItemHscodeDetailGetRequest) GetHscode() string {
    return r.hscode
}

