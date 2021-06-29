package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品搜索人群修改状态 API请求
taobao.simba.serchcrowd.state.batch.update

暂停或启用单品推广搜索人群溢价
*/
type TaobaoSimbaSerchcrowdStateBatchUpdateRequest struct {
    model.Params
    // 被操作者的淘宝昵称
    nick   string
    // 需要修改出价的人群包id,批量传入时用,分割
    adgroupCrowdIds   []int64
    // 推广单元id
    adgroupId   int64
    // 人群状态,0:暂停;1:启用
    state   int64
}

// 初始化TaobaoSimbaSerchcrowdStateBatchUpdateRequest对象
func NewTaobaoSimbaSerchcrowdStateBatchUpdateRequest() *TaobaoSimbaSerchcrowdStateBatchUpdateRequest{
    return &TaobaoSimbaSerchcrowdStateBatchUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.serchcrowd.state.batch.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetNick() string {
    return r.nick
}
// AdgroupCrowdIds Setter
// 需要修改出价的人群包id,批量传入时用,分割
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetAdgroupCrowdIds(adgroupCrowdIds []int64) error {
    r.adgroupCrowdIds = adgroupCrowdIds
    r.Set("adgroup_crowd_ids", adgroupCrowdIds)
    return nil
}

// AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetAdgroupCrowdIds() []int64 {
    return r.adgroupCrowdIds
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// State Setter
// 人群状态,0:暂停;1:启用
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetState(state int64) error {
    r.state = state
    r.Set("state", state)
    return nil
}

// State Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetState() int64 {
    return r.state
}
