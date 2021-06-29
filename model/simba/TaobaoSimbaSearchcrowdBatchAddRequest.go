package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元增加搜索人群 API请求
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

// 初始化TaobaoSimbaSearchcrowdBatchAddRequest对象
func NewTaobaoSimbaSearchcrowdBatchAddRequest() *TaobaoSimbaSearchcrowdBatchAddRequest{
    return &TaobaoSimbaSearchcrowdBatchAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetApiMethodName() string {
    return "taobao.simba.searchcrowd.batch.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSearchcrowdBatchAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetNick() string {
    return r.nick
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSearchcrowdBatchAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// AdgroupTargetingTags Setter
// 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
func (r *TaobaoSimbaSearchcrowdBatchAddRequest) SetAdgroupTargetingTags(adgroupTargetingTags string) error {
    r.adgroupTargetingTags = adgroupTargetingTags
    r.Set("adgroup_targeting_tags", adgroupTargetingTags)
    return nil
}

// AdgroupTargetingTags Getter
func (r TaobaoSimbaSearchcrowdBatchAddRequest) GetAdgroupTargetingTags() string {
    return r.adgroupTargetingTags
}
