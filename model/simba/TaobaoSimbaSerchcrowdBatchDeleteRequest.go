package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品搜索人群批量取消溢价 APIRequest
taobao.simba.serchcrowd.batch.delete

删除单品搜索人群溢价功能
*/
type TaobaoSimbaSerchcrowdBatchDeleteRequest struct {
    model.Params

    // 被操作者的淘宝昵称
    nick   string 

    // 子帐号nick
    subNick   string 

    // 需要删除的人群id
    adgroupCrowdIds   []int64 

}

func NewTaobaoSimbaSerchcrowdBatchDeleteRequest() *TaobaoSimbaSerchcrowdBatchDeleteRequest{
    return &TaobaoSimbaSerchcrowdBatchDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.serchcrowd.batch.delete"
}

func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSerchcrowdBatchDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaSerchcrowdBatchDeleteRequest) SetSubNick(subNick string) error {
    r.subNick = subNick
    r.Set("sub_nick", subNick)
    return nil
}

func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetSubNick() string {
    return r.subNick
}

func (r *TaobaoSimbaSerchcrowdBatchDeleteRequest) SetAdgroupCrowdIds(adgroupCrowdIds []int64) error {
    r.adgroupCrowdIds = adgroupCrowdIds
    r.Set("adgroup_crowd_ids", adgroupCrowdIds)
    return nil
}

func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetAdgroupCrowdIds() []int64 {
    return r.adgroupCrowdIds
}

