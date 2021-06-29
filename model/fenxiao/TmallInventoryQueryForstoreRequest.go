package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品仓库库存 API请求
tmall.inventory.query.forstore

商家查询后端商品仓库库存
*/
type TmallInventoryQueryForstoreRequest struct {
    model.Params
    // 查询列表
    paramList   []InventoryQueryForStoreRequest
}

// 初始化TmallInventoryQueryForstoreRequest对象
func NewTmallInventoryQueryForstoreRequest() *TmallInventoryQueryForstoreRequest{
    return &TmallInventoryQueryForstoreRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallInventoryQueryForstoreRequest) GetApiMethodName() string {
    return "tmall.inventory.query.forstore"
}

// IRequest interface 方法, 获取API参数
func (r TmallInventoryQueryForstoreRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 查询列表
func (r *TmallInventoryQueryForstoreRequest) SetParamList(paramList []InventoryQueryForStoreRequest) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

// ParamList Getter
func (r TmallInventoryQueryForstoreRequest) GetParamList() []InventoryQueryForStoreRequest {
    return r.paramList
}
