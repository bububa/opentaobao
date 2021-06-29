package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品搜索人群批量取消溢价 API请求
taobao.simba.serchcrowd.batch.delete

删除单品搜索人群溢价功能
*/
type TaobaoSimbaSerchcrowdBatchDeleteRequest struct {
    model.Params
    // 被操作者的淘宝昵称
    _nick   string
    // 子帐号nick
    _subNick   string
    // 需要删除的人群id
    _adgroupCrowdIds   []int64
}

// 初始化TaobaoSimbaSerchcrowdBatchDeleteRequest对象
func NewTaobaoSimbaSerchcrowdBatchDeleteRequest() *TaobaoSimbaSerchcrowdBatchDeleteRequest{
    return &TaobaoSimbaSerchcrowdBatchDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.serchcrowd.batch.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdBatchDeleteRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetNick() string {
    return r._nick
}
// SubNick Setter
// 子帐号nick
func (r *TaobaoSimbaSerchcrowdBatchDeleteRequest) SetSubNick(_subNick string) error {
    r._subNick = _subNick
    r.Set("sub_nick", _subNick)
    return nil
}

// SubNick Getter
func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetSubNick() string {
    return r._subNick
}
// AdgroupCrowdIds Setter
// 需要删除的人群id
func (r *TaobaoSimbaSerchcrowdBatchDeleteRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []int64) error {
    r._adgroupCrowdIds = _adgroupCrowdIds
    r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
    return nil
}

// AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdBatchDeleteRequest) GetAdgroupCrowdIds() []int64 {
    return r._adgroupCrowdIds
}
