package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌商品实仓库存/周转效能 APIRequest
alibaba.lst.branddatashare.stockdata.query

品牌商查询授权供应商实仓库存数据
*/
type AlibabaLstBranddatashareStockdataQueryRequest struct {
    model.Params

    // 入参
    param   *BmSupplierStockDataParam 

}

func NewAlibabaLstBranddatashareStockdataQueryRequest() *AlibabaLstBranddatashareStockdataQueryRequest{
    return &AlibabaLstBranddatashareStockdataQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstBranddatashareStockdataQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.branddatashare.stockdata.query"
}

func (r AlibabaLstBranddatashareStockdataQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstBranddatashareStockdataQueryRequest) SetParam(param *BmSupplierStockDataParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaLstBranddatashareStockdataQueryRequest) GetParam() *BmSupplierStockDataParam {
    return r.param
}

