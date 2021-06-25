package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店类目获取接口 APIRequest
taobao.qimen.storecategory.get

商家在ERP中调用该接口，获取门店类目
*/
type TaobaoQimenStorecategoryGetRequest struct {
    model.Params

    // 备注
    remark   string 

}

func NewTaobaoQimenStorecategoryGetRequest() *TaobaoQimenStorecategoryGetRequest{
    return &TaobaoQimenStorecategoryGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStorecategoryGetRequest) GetApiMethodName() string {
    return "taobao.qimen.storecategory.get"
}

func (r TaobaoQimenStorecategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStorecategoryGetRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoQimenStorecategoryGetRequest) GetRemark() string {
    return r.remark
}

