package lstwarehouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstwarehouse"
)

/* 
查询品牌商品实仓库存/周转效能 
alibaba.lst.branddatashare.stockdata.query

品牌商查询授权供应商实仓库存数据
*/
func AlibabaLstBranddatashareStockdataQuery(clt *core.SDKClient, req *lstwarehouse.AlibabaLstBranddatashareStockdataQueryAPIRequest, session string) (*lstwarehouse.AlibabaLstBranddatashareStockdataQueryAPIResponse, error) {
    var resp lstwarehouse.AlibabaLstBranddatashareStockdataQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
