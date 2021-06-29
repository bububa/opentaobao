package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一口价商品上架 API请求
taobao.item.update.listing

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateListingRequest struct {
    model.Params
    // 商品数字ID，该参数必须
    numIid   int64
    // 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
    num   int64
}

// 初始化TaobaoItemUpdateListingRequest对象
func NewTaobaoItemUpdateListingRequest() *TaobaoItemUpdateListingRequest{
    return &TaobaoItemUpdateListingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateListingRequest) GetApiMethodName() string {
    return "taobao.item.update.listing"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateListingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemUpdateListingRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemUpdateListingRequest) GetNumIid() int64 {
    return r.numIid
}
// Num Setter
// 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
func (r *TaobaoItemUpdateListingRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

// Num Getter
func (r TaobaoItemUpdateListingRequest) GetNum() int64 {
    return r.num
}
