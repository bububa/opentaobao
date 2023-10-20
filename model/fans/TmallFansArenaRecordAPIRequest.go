package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfansarenarecordAPIRequest 记录完成擂台的用户 API请求
// tmall.fans.arena.record
//
// 记录完成擂台的用户和完成分数
type TmallfansarenarecordAPIRequest struct {
	model.Params
	// mixnick
	_mixNick string
	// 资金池id
	_cashPoolId int64
	// 用户得分
	_score int64
}

// NewTmallfansarenarecordRequest 初始化TmallfansarenarecordAPIRequest对象
func NewTmallfansarenarecordRequest() *TmallfansarenarecordAPIRequest {
	return &TmallfansarenarecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfansarenarecordAPIRequest) GetApiMethodName() string {
	return "tmall.fans.arena.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfansarenarecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfansarenarecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixNick is MixNick Setter
// mixnick
func (r *TmallfansarenarecordAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TmallfansarenarecordAPIRequest) GetMixNick() string {
	return r._mixNick
}

// SetCashPoolId is CashPoolId Setter
// 资金池id
func (r *TmallfansarenarecordAPIRequest) SetCashPoolId(_cashPoolId int64) error {
	r._cashPoolId = _cashPoolId
	r.Set("cash_pool_id", _cashPoolId)
	return nil
}

// GetCashPoolId CashPoolId Getter
func (r TmallfansarenarecordAPIRequest) GetCashPoolId() int64 {
	return r._cashPoolId
}

// SetScore is Score Setter
// 用户得分
func (r *TmallfansarenarecordAPIRequest) SetScore(_score int64) error {
	r._score = _score
	r.Set("score", _score)
	return nil
}

// GetScore Score Getter
func (r TmallfansarenarecordAPIRequest) GetScore() int64 {
	return r._score
}
