package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
记录完成擂台的用户 API请求
tmall.fans.arena.record

记录完成擂台的用户和完成分数
*/
type TmallFansArenaRecordRequest struct {
    model.Params
    // 资金池id
    _cashPoolId   int64
    // 用户得分
    _score   int64
    // mixnick
    _mixNick   string
}

// 初始化TmallFansArenaRecordRequest对象
func NewTmallFansArenaRecordRequest() *TmallFansArenaRecordRequest{
    return &TmallFansArenaRecordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFansArenaRecordRequest) GetApiMethodName() string {
    return "tmall.fans.arena.record"
}

// IRequest interface 方法, 获取API参数
func (r TmallFansArenaRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CashPoolId Setter
// 资金池id
func (r *TmallFansArenaRecordRequest) SetCashPoolId(_cashPoolId int64) error {
    r._cashPoolId = _cashPoolId
    r.Set("cash_pool_id", _cashPoolId)
    return nil
}

// CashPoolId Getter
func (r TmallFansArenaRecordRequest) GetCashPoolId() int64 {
    return r._cashPoolId
}
// Score Setter
// 用户得分
func (r *TmallFansArenaRecordRequest) SetScore(_score int64) error {
    r._score = _score
    r.Set("score", _score)
    return nil
}

// Score Getter
func (r TmallFansArenaRecordRequest) GetScore() int64 {
    return r._score
}
// MixNick Setter
// mixnick
func (r *TmallFansArenaRecordRequest) SetMixNick(_mixNick string) error {
    r._mixNick = _mixNick
    r.Set("mix_nick", _mixNick)
    return nil
}

// MixNick Getter
func (r TmallFansArenaRecordRequest) GetMixNick() string {
    return r._mixNick
}
