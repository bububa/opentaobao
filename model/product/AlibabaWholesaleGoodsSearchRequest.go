package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批发市场产品搜索 APIRequest
alibaba.wholesale.goods.search

批发市场产品搜索
*/
type AlibabaWholesaleGoodsSearchRequest struct {
    model.Params

    // SearchGoodsOption
    paramSearchGoodsOption   *SearchGoodsOption 

}

func NewAlibabaWholesaleGoodsSearchRequest() *AlibabaWholesaleGoodsSearchRequest{
    return &AlibabaWholesaleGoodsSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWholesaleGoodsSearchRequest) GetApiMethodName() string {
    return "alibaba.wholesale.goods.search"
}

func (r AlibabaWholesaleGoodsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWholesaleGoodsSearchRequest) SetParamSearchGoodsOption(paramSearchGoodsOption *SearchGoodsOption) error {
    r.paramSearchGoodsOption = paramSearchGoodsOption
    r.Set("param_search_goods_option", paramSearchGoodsOption)
    return nil
}

func (r AlibabaWholesaleGoodsSearchRequest) GetParamSearchGoodsOption() *SearchGoodsOption {
    return r.paramSearchGoodsOption
}

