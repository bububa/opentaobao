package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改产品线 API请求
taobao.fenxiao.productcat.update

修改产品线
*/
type TaobaoFenxiaoProductcatUpdateRequest struct {
    model.Params
    // 产品线ID
    _productLineId   int64
    // 产品线名称
    _name   string
    // 最低零售价比例，注意：100.00%，则输入为10000
    _retailLowPercent   int64
    // 最高零售价比例，注意：100.00%，则输入为10000
    _retailHighPercent   int64
    // 代销默认采购价比例，注意：100.00%，则输入为10000
    _agentCostPercent   int64
    // 经销默认采购价比例，注意：100.00%，则输入为10000
    _dealerCostPercent   int64
}

// 初始化TaobaoFenxiaoProductcatUpdateRequest对象
func NewTaobaoFenxiaoProductcatUpdateRequest() *TaobaoFenxiaoProductcatUpdateRequest{
    return &TaobaoFenxiaoProductcatUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcat.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductLineId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetProductLineId(_productLineId int64) error {
    r._productLineId = _productLineId
    r.Set("product_line_id", _productLineId)
    return nil
}

// ProductLineId Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetProductLineId() int64 {
    return r._productLineId
}
// Name Setter
// 产品线名称
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetName() string {
    return r._name
}
// RetailLowPercent Setter
// 最低零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetRetailLowPercent(_retailLowPercent int64) error {
    r._retailLowPercent = _retailLowPercent
    r.Set("retail_low_percent", _retailLowPercent)
    return nil
}

// RetailLowPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetRetailLowPercent() int64 {
    return r._retailLowPercent
}
// RetailHighPercent Setter
// 最高零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetRetailHighPercent(_retailHighPercent int64) error {
    r._retailHighPercent = _retailHighPercent
    r.Set("retail_high_percent", _retailHighPercent)
    return nil
}

// RetailHighPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetRetailHighPercent() int64 {
    return r._retailHighPercent
}
// AgentCostPercent Setter
// 代销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetAgentCostPercent(_agentCostPercent int64) error {
    r._agentCostPercent = _agentCostPercent
    r.Set("agent_cost_percent", _agentCostPercent)
    return nil
}

// AgentCostPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetAgentCostPercent() int64 {
    return r._agentCostPercent
}
// DealerCostPercent Setter
// 经销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetDealerCostPercent(_dealerCostPercent int64) error {
    r._dealerCostPercent = _dealerCostPercent
    r.Set("dealer_cost_percent", _dealerCostPercent)
    return nil
}

// DealerCostPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetDealerCostPercent() int64 {
    return r._dealerCostPercent
}
