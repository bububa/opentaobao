package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批发市场产品搜索 API请求
alibaba.wholesale.goods.search

批发市场产品搜索
*/
type AlibabaWholesaleGoodsSearchAPIRequest struct {
    model.Params
    // SearchGoodsOption
    _paramSearchGoodsOption   *SearchGoodsOption
}

// 初始化AlibabaWholesaleGoodsSearchAPIRequest对象
func NewAlibabaWholesaleGoodsSearchRequest() *AlibabaWholesaleGoodsSearchAPIRequest{
    return &AlibabaWholesaleGoodsSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleGoodsSearchAPIRequest) GetApiMethodName() string {
    return "alibaba.wholesale.goods.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleGoodsSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSearchGoodsOption Setter
// SearchGoodsOption
func (r *AlibabaWholesaleGoodsSearchAPIRequest) SetParamSearchGoodsOption(_paramSearchGoodsOption *SearchGoodsOption) error {
    r._paramSearchGoodsOption = _paramSearchGoodsOption
    r.Set("param_search_goods_option", _paramSearchGoodsOption)
    return nil
}

// ParamSearchGoodsOption Getter
func (r AlibabaWholesaleGoodsSearchAPIRequest) GetParamSearchGoodsOption() *SearchGoodsOption {
    return r._paramSearchGoodsOption
}
