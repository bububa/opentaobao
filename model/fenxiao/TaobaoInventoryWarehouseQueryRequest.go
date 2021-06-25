package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询商家仓信息 APIRequest
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

func NewTaobaoInventoryWarehouseQueryRequest() *TaobaoInventoryWarehouseQueryRequest{
    return &TaobaoInventoryWarehouseQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryWarehouseQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.warehouse.query"
}

func (r TaobaoInventoryWarehouseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryWarehouseQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoInventoryWarehouseQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoInventoryWarehouseQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoInventoryWarehouseQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

