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
    cashPoolId   int64
    // 用户得分
    score   int64
    // mixnick
    mixNick   string
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
func (r *TmallFansArenaRecordRequest) SetCashPoolId(cashPoolId int64) error {
    r.cashPoolId = cashPoolId
    r.Set("cash_pool_id", cashPoolId)
    return nil
}

// CashPoolId Getter
func (r TmallFansArenaRecordRequest) GetCashPoolId() int64 {
    return r.cashPoolId
}
// Score Setter
// 用户得分
func (r *TmallFansArenaRecordRequest) SetScore(score int64) error {
    r.score = score
    r.Set("score", score)
    return nil
}

// Score Getter
func (r TmallFansArenaRecordRequest) GetScore() int64 {
    return r.score
}
// MixNick Setter
// mixnick
func (r *TmallFansArenaRecordRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

// MixNick Getter
func (r TmallFansArenaRecordRequest) GetMixNick() string {
    return r.mixNick
}
