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
type TaobaoSimbaSearchcrowdBatchAddAPIRequest struct {
    model.Params
    // 被操作者的淘宝昵称
    _nick   string
    // 推广单元id
    _adgroupId   int64
    // 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
    _adgroupTargetingTags   string
}

// 初始化TaobaoSimbaSearchcrowdBatchAddAPIRequest对象
func NewTaobaoSimbaSearchcrowdBatchAddRequest() *TaobaoSimbaSearchcrowdBatchAddAPIRequest{
    return &TaobaoSimbaSearchcrowdBatchAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetApiMethodName() string {
    return "taobao.simba.searchcrowd.batch.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSearchcrowdBatchAddAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSearchcrowdBatchAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// AdgroupTargetingTags Setter
// 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
func (r *TaobaoSimbaSearchcrowdBatchAddAPIRequest) SetAdgroupTargetingTags(_adgroupTargetingTags string) error {
    r._adgroupTargetingTags = _adgroupTargetingTags
    r.Set("adgroup_targeting_tags", _adgroupTargetingTags)
    return nil
}

// AdgroupTargetingTags Getter
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetAdgroupTargetingTags() string {
    return r._adgroupTargetingTags
}
