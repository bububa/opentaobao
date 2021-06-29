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
    productLineId   int64
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
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetProductLineId(productLineId int64) error {
    r.productLineId = productLineId
    r.Set("product_line_id", productLineId)
    return nil
}

// ProductLineId Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetProductLineId() int64 {
    return r.productLineId
}
// Name Setter
// 产品线名称
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetName() string {
    return r.name
}
// RetailLowPercent Setter
// 最低零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetRetailLowPercent(retailLowPercent int64) error {
    r.retailLowPercent = retailLowPercent
    r.Set("retail_low_percent", retailLowPercent)
    return nil
}

// RetailLowPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetRetailLowPercent() int64 {
    return r.retailLowPercent
}
// RetailHighPercent Setter
// 最高零售价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetRetailHighPercent(retailHighPercent int64) error {
    r.retailHighPercent = retailHighPercent
    r.Set("retail_high_percent", retailHighPercent)
    return nil
}

// RetailHighPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetRetailHighPercent() int64 {
    return r.retailHighPercent
}
// AgentCostPercent Setter
// 代销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetAgentCostPercent(agentCostPercent int64) error {
    r.agentCostPercent = agentCostPercent
    r.Set("agent_cost_percent", agentCostPercent)
    return nil
}

// AgentCostPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetAgentCostPercent() int64 {
    return r.agentCostPercent
}
// DealerCostPercent Setter
// 经销默认采购价比例，注意：100.00%，则输入为10000
func (r *TaobaoFenxiaoProductcatUpdateRequest) SetDealerCostPercent(dealerCostPercent int64) error {
    r.dealerCostPercent = dealerCostPercent
    r.Set("dealer_cost_percent", dealerCostPercent)
    return nil
}

// DealerCostPercent Getter
func (r TaobaoFenxiaoProductcatUpdateRequest) GetDealerCostPercent() int64 {
    return r.dealerCostPercent
}
