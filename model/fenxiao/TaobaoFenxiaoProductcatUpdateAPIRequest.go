package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductcatupdateAPIRequest 修改产品线 API请求
// taobao.fenxiao.productcat.update
//
// 修改产品线
type TaobaofenxiaoproductcatupdateAPIRequest struct {
	model.Params
	// 产品线名称
	_name string
	// 产品线ID
	_productLineId int64
	// 最低零售价比例，注意：100.00%，则输入为10000
	_retailLowPercent int64
	// 最高零售价比例，注意：100.00%，则输入为10000
	_retailHighPercent int64
	// 代销默认采购价比例，注意：100.00%，则输入为10000
	_agentCostPercent int64
	// 经销默认采购价比例，注意：100.00%，则输入为10000
	_dealerCostPercent int64
}

// NewTaobaofenxiaoproductcatupdateRequest 初始化TaobaofenxiaoproductcatupdateAPIRequest对象
func NewTaobaofenxiaoproductcatupdateRequest() *TaobaofenxiaoproductcatupdateAPIRequest {
	return &TaobaofenxiaoproductcatupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcat.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 产品线名称
func (r *TaobaofenxiaoproductcatupdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetName() string {
	return r._name
}

// SetProductLineId is ProductLineId Setter
// 产品线ID
func (r *TaobaofenxiaoproductcatupdateAPIRequest) SetProductLineId(_productLineId int64) error {
	r._productLineId = _productLineId
	r.Set("product_line_id", _productLineId)
	return nil
}

// GetProductLineId ProductLineId Getter
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetProductLineId() int64 {
	return r._productLineId
}

// SetRetailLowPercent is RetailLowPercent Setter
// 最低零售价比例，注意：100.00%，则输入为10000
func (r *TaobaofenxiaoproductcatupdateAPIRequest) SetRetailLowPercent(_retailLowPercent int64) error {
	r._retailLowPercent = _retailLowPercent
	r.Set("retail_low_percent", _retailLowPercent)
	return nil
}

// GetRetailLowPercent RetailLowPercent Getter
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetRetailLowPercent() int64 {
	return r._retailLowPercent
}

// SetRetailHighPercent is RetailHighPercent Setter
// 最高零售价比例，注意：100.00%，则输入为10000
func (r *TaobaofenxiaoproductcatupdateAPIRequest) SetRetailHighPercent(_retailHighPercent int64) error {
	r._retailHighPercent = _retailHighPercent
	r.Set("retail_high_percent", _retailHighPercent)
	return nil
}

// GetRetailHighPercent RetailHighPercent Getter
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetRetailHighPercent() int64 {
	return r._retailHighPercent
}

// SetAgentCostPercent is AgentCostPercent Setter
// 代销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaofenxiaoproductcatupdateAPIRequest) SetAgentCostPercent(_agentCostPercent int64) error {
	r._agentCostPercent = _agentCostPercent
	r.Set("agent_cost_percent", _agentCostPercent)
	return nil
}

// GetAgentCostPercent AgentCostPercent Getter
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetAgentCostPercent() int64 {
	return r._agentCostPercent
}

// SetDealerCostPercent is DealerCostPercent Setter
// 经销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaofenxiaoproductcatupdateAPIRequest) SetDealerCostPercent(_dealerCostPercent int64) error {
	r._dealerCostPercent = _dealerCostPercent
	r.Set("dealer_cost_percent", _dealerCostPercent)
	return nil
}

// GetDealerCostPercent DealerCostPercent Getter
func (r TaobaofenxiaoproductcatupdateAPIRequest) GetDealerCostPercent() int64 {
	return r._dealerCostPercent
}
