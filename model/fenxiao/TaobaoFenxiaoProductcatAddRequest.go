package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增产品线 API请求
taobao.fenxiao.productcat.add

新增产品线
*/
type TaobaoFenxiaoProductcatAddAPIRequest struct {
    model.Params
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

// 初始化TaobaoFenxiaoProductcatAddAPIRequest对象
func NewTaobaoFenxiaoProductcatAddRequest() *TaobaoFenxiaoProductcatAddAPIRequest{
    return &TaobaoFenxiaoProductcatAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcat.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 产品线名称
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetName() string {
    return r._name
}
// RetailLowPercent Setter
// 最低零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetRetailLowPercent(_retailLowPercent int64) error {
    r._retailLowPercent = _retailLowPercent
    r.Set("retail_low_percent", _retailLowPercent)
    return nil
}

// RetailLowPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetRetailLowPercent() int64 {
    return r._retailLowPercent
}
// RetailHighPercent Setter
// 最高零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetRetailHighPercent(_retailHighPercent int64) error {
    r._retailHighPercent = _retailHighPercent
    r.Set("retail_high_percent", _retailHighPercent)
    return nil
}

// RetailHighPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetRetailHighPercent() int64 {
    return r._retailHighPercent
}
// AgentCostPercent Setter
// 代销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetAgentCostPercent(_agentCostPercent int64) error {
    r._agentCostPercent = _agentCostPercent
    r.Set("agent_cost_percent", _agentCostPercent)
    return nil
}

// AgentCostPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetAgentCostPercent() int64 {
    return r._agentCostPercent
}
// DealerCostPercent Setter
// 经销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddAPIRequest) SetDealerCostPercent(_dealerCostPercent int64) error {
    r._dealerCostPercent = _dealerCostPercent
    r.Set("dealer_cost_percent", _dealerCostPercent)
    return nil
}

// DealerCostPercent Getter
func (r TaobaoFenxiaoProductcatAddAPIRequest) GetDealerCostPercent() int64 {
    return r._dealerCostPercent
}
