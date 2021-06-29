package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
记录完成擂台的用户 APIRequest
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

func NewTmallFansArenaRecordRequest() *TmallFansArenaRecordRequest{
    return &TmallFansArenaRecordRequest{
        Params: model.NewParams(),
    }
}

func (r TmallFansArenaRecordRequest) GetApiMethodName() string {
    return "tmall.fans.arena.record"
}

func (r TmallFansArenaRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallFansArenaRecordRequest) SetCashPoolId(cashPoolId int64) error {
    r.cashPoolId = cashPoolId
    r.Set("cash_pool_id", cashPoolId)
    return nil
}

func (r TmallFansArenaRecordRequest) GetCashPoolId() int64 {
    return r.cashPoolId
}

func (r *TmallFansArenaRecordRequest) SetScore(score int64) error {
    r.score = score
    r.Set("score", score)
    return nil
}

func (r TmallFansArenaRecordRequest) GetScore() int64 {
    return r.score
}

func (r *TmallFansArenaRecordRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

func (r TmallFansArenaRecordRequest) GetMixNick() string {
    return r.mixNick
}

