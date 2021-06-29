package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌商品实仓库存/周转效能 API请求
alibaba.lst.branddatashare.stockdata.query

品牌商查询授权供应商实仓库存数据
*/
type AlibabaLstBranddatashareStockdataQueryRequest struct {
    model.Params
    // 入参
    param   *BmSupplierStockDataParam
}

// 初始化AlibabaLstBranddatashareStockdataQueryRequest对象
func NewAlibabaLstBranddatashareStockdataQueryRequest() *AlibabaLstBranddatashareStockdataQueryRequest{
    return &AlibabaLstBranddatashareStockdataQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstBranddatashareStockdataQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.branddatashare.stockdata.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstBranddatashareStockdataQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaLstBranddatashareStockdataQueryRequest) SetParam(param *BmSupplierStockDataParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaLstBranddatashareStockdataQueryRequest) GetParam() *BmSupplierStockDataParam {
    return r.param
}
