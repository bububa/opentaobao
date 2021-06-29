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
    _nick   string
    // 需要修改出价的人群包id,批量传入时用,分割
    _adgroupCrowdIds   []int64
    // 推广单元id
    _adgroupId   int64
    // 人群状态,0:暂停;1:启用
    _state   int64
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
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetNick() string {
    return r._nick
}
// AdgroupCrowdIds Setter
// 需要修改出价的人群包id,批量传入时用,分割
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []int64) error {
    r._adgroupCrowdIds = _adgroupCrowdIds
    r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
    return nil
}

// AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetAdgroupCrowdIds() []int64 {
    return r._adgroupCrowdIds
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// State Setter
// 人群状态,0:暂停;1:启用
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateRequest) SetState(_state int64) error {
    r._state = _state
    r.Set("state", _state)
    return nil
}

// State Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateRequest) GetState() int64 {
    return r._state
}
