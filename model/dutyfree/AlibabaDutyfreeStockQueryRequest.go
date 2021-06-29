package dutyfree

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外库存查询接口 API请求
alibaba.dutyfree.stock.query

对外部服务提供库存查询接口
*/
type AlibabaDutyfreeStockQueryRequest struct {
    model.Params
    // 条形码
    _barCode   string
}

// 初始化AlibabaDutyfreeStockQueryRequest对象
func NewAlibabaDutyfreeStockQueryRequest() *AlibabaDutyfreeStockQueryRequest{
    return &AlibabaDutyfreeStockQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDutyfreeStockQueryRequest) GetApiMethodName() string {
    return "alibaba.dutyfree.stock.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDutyfreeStockQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BarCode Setter
// 条形码
func (r *AlibabaDutyfreeStockQueryRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r AlibabaDutyfreeStockQueryRequest) GetBarCode() string {
    return r._barCode
}
