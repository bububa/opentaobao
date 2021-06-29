package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过hscode获取计量单位 API请求
tmall.item.hscode.detail.get

通过hscode获取计量单位和销售单位
*/
type TmallItemHscodeDetailGetRequest struct {
    model.Params
    // hscode
    hscode   string
}

// 初始化TmallItemHscodeDetailGetRequest对象
func NewTmallItemHscodeDetailGetRequest() *TmallItemHscodeDetailGetRequest{
    return &TmallItemHscodeDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemHscodeDetailGetRequest) GetApiMethodName() string {
    return "tmall.item.hscode.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemHscodeDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hscode Setter
// hscode
func (r *TmallItemHscodeDetailGetRequest) SetHscode(hscode string) error {
    r.hscode = hscode
    r.Set("hscode", hscode)
    return nil
}

// Hscode Getter
func (r TmallItemHscodeDetailGetRequest) GetHscode() string {
    return r.hscode
}
