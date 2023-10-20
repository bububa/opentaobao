package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatAddAPIRequest 新增产品线 API请求
// taobao.fenxiao.productcat.add
//
// 新增产品线
type TaobaoFenxiaoProductcatAddAPIRequest struct {
	model.Params
	// 产品线名称
	_name string
	// 最低零售价比例，注意：100.00%，则输入为10000
	_retailLowPercent int64
	// 最高零售价比例，注意：100.00%，则输入为10000
	_retailHighPercent int64
	// 代销默认采购价比例，注意：100.00%，则输入为10000
	_agentCostPercent int64
	// 经销默认采购价比例，注意：100.00%，则输入为10000
	_dealerCostPercent int64
}

// NewTaobaoFenxiaoProductcatAddRequest 初始化TaobaoFenxiaoProductcatAddAPIRequest对象
func NewTaobaoFenxiaoProductcatAddRequest() *TaobaoFenxiaoProductcatAddAPIRequest {
	return &TaobaoFenxiaoProductcatAddAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductcatAddAPIRequest) Reset() {
	r._name = ""
	r._retailLowPercent = 0
	r._retailHighPercent = 0
	r._agentCostPercent = 0
	r._dealerCostPercent = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcat.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 产品线名称
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetName() string {
	return r._name
}

// SetRetailLowPercent is RetailLowPercent Setter
// 最低零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetRetailLowPercent(_retailLowPercent int64) error {
	r._retailLowPercent = _retailLowPercent
	r.Set("retail_low_percent", _retailLowPercent)
	return nil
}

// GetRetailLowPercent RetailLowPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetRetailLowPercent() int64 {
	return r._retailLowPercent
}

// SetRetailHighPercent is RetailHighPercent Setter
// 最高零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetRetailHighPercent(_retailHighPercent int64) error {
	r._retailHighPercent = _retailHighPercent
	r.Set("retail_high_percent", _retailHighPercent)
	return nil
}

// GetRetailHighPercent RetailHighPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetRetailHighPercent() int64 {
	return r._retailHighPercent
}

// SetAgentCostPercent is AgentCostPercent Setter
// 代销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetAgentCostPercent(_agentCostPercent int64) error {
	r._agentCostPercent = _agentCostPercent
	r.Set("agent_cost_percent", _agentCostPercent)
	return nil
}

// GetAgentCostPercent AgentCostPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetAgentCostPercent() int64 {
	return r._agentCostPercent
}

// SetDealerCostPercent is DealerCostPercent Setter
// 经销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetDealerCostPercent(_dealerCostPercent int64) error {
	r._dealerCostPercent = _dealerCostPercent
	r.Set("dealer_cost_percent", _dealerCostPercent)
	return nil
}

// GetDealerCostPercent DealerCostPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetDealerCostPercent() int64 {
	return r._dealerCostPercent
}

var poolTaobaoFenxiaoProductcatAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductcatAddRequest()
	},
}

// GetTaobaoFenxiaoProductcatAddRequest 从 sync.Pool 获取 TaobaoFenxiaoProductcatAddAPIRequest
func GetTaobaoFenxiaoProductcatAddAPIRequest() *TaobaoFenxiaoProductcatAddAPIRequest {
	return poolTaobaoFenxiaoProductcatAddAPIRequest.Get().(*TaobaoFenxiaoProductcatAddAPIRequest)
}

// ReleaseTaobaoFenxiaoProductcatAddAPIRequest 将 TaobaoFenxiaoProductcatAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductcatAddAPIRequest(v *TaobaoFenxiaoProductcatAddAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductcatAddAPIRequest.Put(v)
}
