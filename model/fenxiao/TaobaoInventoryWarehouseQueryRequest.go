package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询商家仓信息 API请求
taobao.inventory.warehouse.query

分页查询商家仓信息
*/
type TaobaoInventoryWarehouseQueryRequest struct {
    model.Params
    // 页码
    _pageNo   int64
    // 页大小
    _pageSize   int64
}

// 初始化TaobaoInventoryWarehouseQueryRequest对象
func NewTaobaoInventoryWarehouseQueryRequest() *TaobaoInventoryWarehouseQueryRequest{
    return &TaobaoInventoryWarehouseQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryWarehouseQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.warehouse.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryWarehouseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 页码
func (r *TaobaoInventoryWarehouseQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoInventoryWarehouseQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 页大小
func (r *TaobaoInventoryWarehouseQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoInventoryWarehouseQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
