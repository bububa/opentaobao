package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
批发市场产品搜索 
alibaba.wholesale.goods.search

批发市场产品搜索
*/
func AlibabaWholesaleGoodsSearch(clt *core.SDKClient, req *product.AlibabaWholesaleGoodsSearchAPIRequest, session string) (*product.AlibabaWholesaleGoodsSearchAPIResponse, error) {
    var resp product.AlibabaWholesaleGoodsSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
