package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rellocate APIRequest
taobao.omniorder.store.reallocate

门店发货提供改派接口
*/
type TaobaoOmniorderStoreReallocateRequest struct {
    model.Params

    // 主订单号
    mainOrderId   int64 

    // 子订单号
    subOrderIds   []int64 

    // 门店Id
    storeId   int64 

    // 电商仓code
    warehouseCode   string 

}

func NewTaobaoOmniorderStoreReallocateRequest() *TaobaoOmniorderStoreReallocateRequest{
    return &TaobaoOmniorderStoreReallocateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreReallocateRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.reallocate"
}

func (r TaobaoOmniorderStoreReallocateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreReallocateRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoOmniorderStoreReallocateRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoOmniorderStoreReallocateRequest) SetSubOrderIds(subOrderIds []int64) error {
    r.subOrderIds = subOrderIds
    r.Set("sub_order_ids", subOrderIds)
    return nil
}

func (r TaobaoOmniorderStoreReallocateRequest) GetSubOrderIds() []int64 {
    return r.subOrderIds
}

func (r *TaobaoOmniorderStoreReallocateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderStoreReallocateRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoOmniorderStoreReallocateRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoOmniorderStoreReallocateRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

