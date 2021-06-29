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
    pageNo   int64
    // 页大小
    pageSize   int64
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
func (r *TaobaoInventoryWarehouseQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoInventoryWarehouseQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 页大小
func (r *TaobaoInventoryWarehouseQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoInventoryWarehouseQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
