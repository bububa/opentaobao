package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest 根据交易单号判断是否为菜鸟发货订单 API请求
// cainiao.cboss.workplatform.logistics.iscainiaoorder
//
// 根据交易单号判断是否为菜鸟发货订单
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest struct {
	model.Params
	// 交易单号
	_tradeId string
}

// NewCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest 初始化CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest对象
func NewCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest() *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest {
	return &CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) Reset() {
	r._tradeId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) GetApiMethodName() string {
	return "cainiao.cboss.workplatform.logistics.iscainiaoorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// 交易单号
func (r *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) GetTradeId() string {
	return r._tradeId
}

var poolCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest()
	},
}

// GetCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest 从 sync.Pool 获取 CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest
func GetCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest() *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest {
	return poolCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest.Get().(*CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest)
}

// ReleaseCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest 将 CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest 放入 sync.Pool
func ReleaseCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest(v *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) {
	v.Reset()
	poolCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest.Put(v)
}
