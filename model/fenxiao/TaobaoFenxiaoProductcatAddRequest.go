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
type TaobaoFenxiaoProductcatAddRequest struct {
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

// 初始化TaobaoFenxiaoProductcatAddRequest对象
func NewTaobaoFenxiaoProductcatAddRequest() *TaobaoFenxiaoProductcatAddRequest{
    return &TaobaoFenxiaoProductcatAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcat.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 产品线名称
func (r *TaobaoFenxiaoProductcatAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetName() string {
    return r._name
}
// RetailLowPercent Setter
// 最低零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetRetailLowPercent(_retailLowPercent int64) error {
    r._retailLowPercent = _retailLowPercent
    r.Set("retail_low_percent", _retailLowPercent)
    return nil
}

// RetailLowPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetRetailLowPercent() int64 {
    return r._retailLowPercent
}
// RetailHighPercent Setter
// 最高零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetRetailHighPercent(_retailHighPercent int64) error {
    r._retailHighPercent = _retailHighPercent
    r.Set("retail_high_percent", _retailHighPercent)
    return nil
}

// RetailHighPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetRetailHighPercent() int64 {
    return r._retailHighPercent
}
// AgentCostPercent Setter
// 代销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetAgentCostPercent(_agentCostPercent int64) error {
    r._agentCostPercent = _agentCostPercent
    r.Set("agent_cost_percent", _agentCostPercent)
    return nil
}

// AgentCostPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetAgentCostPercent() int64 {
    return r._agentCostPercent
}
// DealerCostPercent Setter
// 经销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetDealerCostPercent(_dealerCostPercent int64) error {
    r._dealerCostPercent = _dealerCostPercent
    r.Set("dealer_cost_percent", _dealerCostPercent)
    return nil
}

// DealerCostPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetDealerCostPercent() int64 {
    return r._dealerCostPercent
}
