package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品推广搜索人群修改溢价 APIRequest
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

func NewTaobaoSimbaSerchcrowdPriceBatchUpdateRequest() *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest{
    return &TaobaoSimbaSerchcrowdPriceBatchUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.serchcrowd.price.batch.update"
}

func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetSubNick(subNick string) error {
    r.subNick = subNick
    r.Set("sub_nick", subNick)
    return nil
}

func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetSubNick() string {
    return r.subNick
}

func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetAdgroupCrowdIds(adgroupCrowdIds []int64) error {
    r.adgroupCrowdIds = adgroupCrowdIds
    r.Set("adgroup_crowd_ids", adgroupCrowdIds)
    return nil
}

func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetAdgroupCrowdIds() []int64 {
    return r.adgroupCrowdIds
}

func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) SetDiscount(discount int64) error {
    r.discount = discount
    r.Set("discount", discount)
    return nil
}

func (r TaobaoSimbaSerchcrowdPriceBatchUpdateRequest) GetDiscount() int64 {
    return r.discount
}

