package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店类目获取接口 API请求
taobao.qimen.storecategory.get

商家在ERP中调用该接口，获取门店类目
*/
type TaobaoQimenStorecategoryGetRequest struct {
    model.Params
    // 备注
    _remark   string
}

// 初始化TaobaoQimenStorecategoryGetRequest对象
func NewTaobaoQimenStorecategoryGetRequest() *TaobaoQimenStorecategoryGetRequest{
    return &TaobaoQimenStorecategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStorecategoryGetRequest) GetApiMethodName() string {
    return "taobao.qimen.storecategory.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStorecategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Remark Setter
// 备注
func (r *TaobaoQimenStorecategoryGetRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenStorecategoryGetRequest) GetRemark() string {
    return r._remark
}
