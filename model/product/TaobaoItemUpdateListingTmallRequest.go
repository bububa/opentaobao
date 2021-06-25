package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
taobao.item.update.listing天猫分流 APIRequest
taobao.item.update.listing.tmall

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateListingTmallRequest struct {
    model.Params

    // 商品数字ID，该参数必须
    numIid   int64 

    // 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
    num   int64 

}

func NewTaobaoItemUpdateListingTmallRequest() *TaobaoItemUpdateListingTmallRequest{
    return &TaobaoItemUpdateListingTmallRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemUpdateListingTmallRequest) GetApiMethodName() string {
    return "taobao.item.update.listing.tmall"
}

func (r TaobaoItemUpdateListingTmallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemUpdateListingTmallRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemUpdateListingTmallRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoItemUpdateListingTmallRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

func (r TaobaoItemUpdateListingTmallRequest) GetNum() int64 {
    return r.num
}

