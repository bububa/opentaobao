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
    name   string
    // 最低零售价比例，注意：100.00%，则输入为10000
    retailLowPercent   int64
    // 最高零售价比例，注意：100.00%，则输入为10000
    retailHighPercent   int64
    // 代销默认采购价比例，注意：100.00%，则输入为10000
    agentCostPercent   int64
    // 经销默认采购价比例，注意：100.00%，则输入为10000
    dealerCostPercent   int64
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
func (r *TaobaoFenxiaoProductcatAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetName() string {
    return r.name
}
// RetailLowPercent Setter
// 最低零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetRetailLowPercent(retailLowPercent int64) error {
    r.retailLowPercent = retailLowPercent
    r.Set("retail_low_percent", retailLowPercent)
    return nil
}

// RetailLowPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetRetailLowPercent() int64 {
    return r.retailLowPercent
}
// RetailHighPercent Setter
// 最高零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetRetailHighPercent(retailHighPercent int64) error {
    r.retailHighPercent = retailHighPercent
    r.Set("retail_high_percent", retailHighPercent)
    return nil
}

// RetailHighPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetRetailHighPercent() int64 {
    return r.retailHighPercent
}
// AgentCostPercent Setter
// 代销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetAgentCostPercent(agentCostPercent int64) error {
    r.agentCostPercent = agentCostPercent
    r.Set("agent_cost_percent", agentCostPercent)
    return nil
}

// AgentCostPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetAgentCostPercent() int64 {
    return r.agentCostPercent
}
// DealerCostPercent Setter
// 经销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatAddRequest) SetDealerCostPercent(dealerCostPercent int64) error {
    r.dealerCostPercent = dealerCostPercent
    r.Set("dealer_cost_percent", dealerCostPercent)
    return nil
}

// DealerCostPercent Getter
func (r TaobaoFenxiaoProductcatAddRequest) GetDealerCostPercent() int64 {
    return r.dealerCostPercent
}
