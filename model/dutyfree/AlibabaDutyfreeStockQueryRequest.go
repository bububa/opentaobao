package dutyfree

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外库存查询接口 APIRequest
alibaba.dutyfree.stock.query

对外部服务提供库存查询接口
*/
type AlibabaDutyfreeStockQueryRequest struct {
    model.Params

    // 条形码
    barCode   string 

}

func NewAlibabaDutyfreeStockQueryRequest() *AlibabaDutyfreeStockQueryRequest{
    return &AlibabaDutyfreeStockQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDutyfreeStockQueryRequest) GetApiMethodName() string {
    return "alibaba.dutyfree.stock.query"
}

func (r AlibabaDutyfreeStockQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDutyfreeStockQueryRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

func (r AlibabaDutyfreeStockQueryRequest) GetBarCode() string {
    return r.barCode
}

