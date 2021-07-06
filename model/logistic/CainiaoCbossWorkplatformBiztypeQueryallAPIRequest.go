package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformBiztypeQueryallAPIRequest 菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 API请求
// cainiao.cboss.workplatform.biztype.queryall
//
// 菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型。 目前调用者ISV
type CainiaoCbossWorkplatformBiztypeQueryallAPIRequest struct {
	model.Params
	// tradeId
	_tradeId string
	// level
	_level int64
}

// NewCainiaoCbossWorkplatformBiztypeQueryallRequest 初始化CainiaoCbossWorkplatformBiztypeQueryallAPIRequest对象
func NewCainiaoCbossWorkplatformBiztypeQueryallRequest() *CainiaoCbossWorkplatformBiztypeQueryallAPIRequest {
	return &CainiaoCbossWorkplatformBiztypeQueryallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformBiztypeQueryallAPIRequest) GetApiMethodName() string {
	return "cainiao.cboss.workplatform.biztype.queryall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformBiztypeQueryallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeId is TradeId Setter
// tradeId
func (r *CainiaoCbossWorkplatformBiztypeQueryallAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r CainiaoCbossWorkplatformBiztypeQueryallAPIRequest) GetTradeId() string {
	return r._tradeId
}

// SetLevel is Level Setter
// level
func (r *CainiaoCbossWorkplatformBiztypeQueryallAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r CainiaoCbossWorkplatformBiztypeQueryallAPIRequest) GetLevel() int64 {
	return r._level
}
