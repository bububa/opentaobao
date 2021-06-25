package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
查询阿里巴巴批发市场商品详情 
alibaba.wholesale.goods.get

查询阿里巴巴批发市场商品详情
*/
func AlibabaWholesaleGoodsGet(clt *core.SDKClient, req *product.AlibabaWholesaleGoodsGetRequest, session string) (*product.AlibabaWholesaleGoodsGetResponse, error) {
    var resp product.AlibabaWholesaleGoodsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
