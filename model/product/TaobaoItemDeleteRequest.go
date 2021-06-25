package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单条商品 APIRequest
taobao.item.delete

删除单条商品
*/
type TaobaoItemDeleteRequest struct {
    model.Params

    // 商品数字ID，该参数必须
    numIid   int64 

}

func NewTaobaoItemDeleteRequest() *TaobaoItemDeleteRequest{
    return &TaobaoItemDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemDeleteRequest) GetApiMethodName() string {
    return "taobao.item.delete"
}

func (r TaobaoItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemDeleteRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemDeleteRequest) GetNumIid() int64 {
    return r.numIid
}

