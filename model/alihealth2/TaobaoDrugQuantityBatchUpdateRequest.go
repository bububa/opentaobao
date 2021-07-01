package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量同步库存接口 API请求
taobao.drug.quantity.batch.update

商家通过top接口可以批量修改商品库存
*/
type TaobaoDrugQuantityBatchUpdateAPIRequest struct {
    model.Params
    // 外部店铺ID
    _outStoreId   string
    // 商品ID和库存
    _outItemIdQuantityMap   string
}

// 初始化TaobaoDrugQuantityBatchUpdateAPIRequest对象
func NewTaobaoDrugQuantityBatchUpdateRequest() *TaobaoDrugQuantityBatchUpdateAPIRequest{
    return &TaobaoDrugQuantityBatchUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.drug.quantity.batch.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugQuantityBatchUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
    r._outStoreId = _outStoreId
    r.Set("out_store_id", _outStoreId)
    return nil
}

// OutStoreId Getter
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetOutStoreId() string {
    return r._outStoreId
}
// OutItemIdQuantityMap Setter
// 商品ID和库存
func (r *TaobaoDrugQuantityBatchUpdateAPIRequest) SetOutItemIdQuantityMap(_outItemIdQuantityMap string) error {
    r._outItemIdQuantityMap = _outItemIdQuantityMap
    r.Set("out_item_id_quantity_map", _outItemIdQuantityMap)
    return nil
}

// OutItemIdQuantityMap Getter
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetOutItemIdQuantityMap() string {
    return r._outItemIdQuantityMap
}
