package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品下架 APIRequest
taobao.item.update.delisting

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateDelistingRequest struct {
    model.Params

    // 商品数字ID，该参数必须
    numIid   int64 

}

func NewTaobaoItemUpdateDelistingRequest() *TaobaoItemUpdateDelistingRequest{
    return &TaobaoItemUpdateDelistingRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemUpdateDelistingRequest) GetApiMethodName() string {
    return "taobao.item.update.delisting"
}

func (r TaobaoItemUpdateDelistingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemUpdateDelistingRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemUpdateDelistingRequest) GetNumIid() int64 {
    return r.numIid
}

