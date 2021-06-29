package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁商品四级车型信息 API请求
tmall.car.lease.itemcarinfo

整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传
*/
type TmallCarLeaseItemcarinfoRequest struct {
    model.Params
    // 商品id
    itemId   int64
}

// 初始化TmallCarLeaseItemcarinfoRequest对象
func NewTmallCarLeaseItemcarinfoRequest() *TmallCarLeaseItemcarinfoRequest{
    return &TmallCarLeaseItemcarinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseItemcarinfoRequest) GetApiMethodName() string {
    return "tmall.car.lease.itemcarinfo"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseItemcarinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallCarLeaseItemcarinfoRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallCarLeaseItemcarinfoRequest) GetItemId() int64 {
    return r.itemId
}
