package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
taobao.item.update.delisting.tmall APIRequest
taobao.item.update.delisting.tmall

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateDelistingTmallRequest struct {
    model.Params

    // 商品数字ID，该参数必须
    numIid   int64 

}

func NewTaobaoItemUpdateDelistingTmallRequest() *TaobaoItemUpdateDelistingTmallRequest{
    return &TaobaoItemUpdateDelistingTmallRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemUpdateDelistingTmallRequest) GetApiMethodName() string {
    return "taobao.item.update.delisting.tmall"
}

func (r TaobaoItemUpdateDelistingTmallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemUpdateDelistingTmallRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemUpdateDelistingTmallRequest) GetNumIid() int64 {
    return r.numIid
}

