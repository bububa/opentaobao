package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品推广搜索人群修改溢价 API请求
taobao.simba.serchcrowd.price.batch.update

单品推广搜索人群修改溢价, 不支持跨推广单元修改
*/
type TaobaoSimbaSerchcrowdPriceBatchUpdateRequest struct {
    model.Params
    // 被操作者的淘宝昵称
    nick   string
    // 子帐号nick
    subNick   string
    // 需要修改出价的人群包id,批量传入的时候用,分割
    adgroupCrowdIds   []int64
    // 推广单元id
    adgroupId   int64
    // 人群溢价比例，溢价范围[5,300]
    discount   int64
}

// 初始化TaobaoSimbaSerchcrowdPriceBatchUpdateRequest对象
func NewTaobaoSimbaSerchcrowdPriceBatchUpdateRequest() *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest{
    return &TaobaoSimbaSerchcrowdPriceBatchUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.serchcrowd.price.batch.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetNick() string {
    return r.nick
}
// SubNick Setter
// 子帐号nick
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetSubNick(subNick string) error {
    r.subNick = subNick
    r.Set("sub_nick", subNick)
    return nil
}

// SubNick Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetSubNick() string {
    return r.subNick
}
// AdgroupCrowdIds Setter
// 需要修改出价的人群包id,批量传入的时候用,分割
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetAdgroupCrowdIds(adgroupCrowdIds []int64) error {
    r.adgroupCrowdIds = adgroupCrowdIds
    r.Set("adgroup_crowd_ids", adgroupCrowdIds)
    return nil
}

// AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetAdgroupCrowdIds() []int64 {
    return r.adgroupCrowdIds
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// Discount Setter
// 人群溢价比例，溢价范围[5,300]
func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetDiscount(discount int64) error {
    r.discount = discount
    r.Set("discount", discount)
    return nil
}

// Discount Getter
func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetDiscount() int64 {
    return r.discount
}
