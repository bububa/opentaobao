package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元增加搜索人群 APIRequest
taobao.simba.searchcrowd.batch.add

推广单元新增搜索人群
*/
type TaobaoSimbaSearchcrowdBatchAddRequest struct {
    model.Params

    // 被操作者的淘宝昵称
    nick   string 

    // 推广单元id
    adgroupId   int64 

    // 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
    adgroupTargetingTags   string 

}

func NewTaobaoSimbaSearchcrowdBatchAddRequest() *TaobaoSimbaSearchcrowdBatchAddRequest{
    return &TaobaoSimbaSearchcrowdBatchAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetApiMethodName() string {
    return "taobao.simba.searchcrowd.batch.add"
}

func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSearchcrowdBatchAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaSearchcrowdBatchAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaSearchcrowdBatchAddRequest) SetAdgroupTargetingTags(adgroupTargetingTags string) error {
    r.adgroupTargetingTags = adgroupTargetingTags
    r.Set("adgroup_targeting_tags", adgroupTargetingTags)
    return nil
}

func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetAdgroupTargetingTags() string {
    return r.adgroupTargetingTags
}

