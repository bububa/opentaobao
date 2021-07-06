package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallFansArenaRecordAPIRequest 记录完成擂台的用户 API请求
// tmall.fans.arena.record
//
// 记录完成擂台的用户和完成分数
type TmallFansArenaRecordAPIRequest struct {
	model.Params
	// mixnick
	_mixNick string
	// 资金池id
	_cashPoolId int64
	// 用户得分
	_score int64
}

// NewTmallFansArenaRecordRequest 初始化TmallFansArenaRecordAPIRequest对象
func NewTmallFansArenaRecordRequest() *TmallFansArenaRecordAPIRequest {
	return &TmallFansArenaRecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFansArenaRecordAPIRequest) GetApiMethodName() string {
	return "tmall.fans.arena.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFansArenaRecordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMixNick is MixNick Setter
// mixnick
func (r *TmallFansArenaRecordAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TmallFansArenaRecordAPIRequest) GetMixNick() string {
	return r._mixNick
}

// SetCashPoolId is CashPoolId Setter
// 资金池id
func (r *TmallFansArenaRecordAPIRequest) SetCashPoolId(_cashPoolId int64) error {
	r._cashPoolId = _cashPoolId
	r.Set("cash_pool_id", _cashPoolId)
	return nil
}

// GetCashPoolId CashPoolId Getter
func (r TmallFansArenaRecordAPIRequest) GetCashPoolId() int64 {
	return r._cashPoolId
}

// SetScore is Score Setter
// 用户得分
func (r *TmallFansArenaRecordAPIRequest) SetScore(_score int64) error {
	r._score = _score
	r.Set("score", _score)
	return nil
}

// GetScore Score Getter
func (r TmallFansArenaRecordAPIRequest) GetScore() int64 {
	return r._score
}
