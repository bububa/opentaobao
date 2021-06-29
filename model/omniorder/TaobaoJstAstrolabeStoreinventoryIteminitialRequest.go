package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存初始化接口 API请求
taobao.jst.astrolabe.storeinventory.iteminitial

ERP调用奇门的接口，对门店的库存进行初始化
*/
type TaobaoJstAstrolabeStoreinventoryIteminitialRequest struct {
    model.Params
    // 门店列表
    _stores   []Store
    // 操作时间
    _operationTime   string
}

// 初始化TaobaoJstAstrolabeStoreinventoryIteminitialRequest对象
func NewTaobaoJstAstrolabeStoreinventoryIteminitialRequest() *TaobaoJstAstrolabeStoreinventoryIteminitialRequest{
    return &TaobaoJstAstrolabeStoreinventoryIteminitialRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.iteminitial"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryIteminitialRequest) SetStores(_stores []Store) error {
    r._stores = _stores
    r.Set("stores", _stores)
    return nil
}

// Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetStores() []Store {
    return r._stores
}
// OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryIteminitialRequest) SetOperationTime(_operationTime string) error {
    r._operationTime = _operationTime
    r.Set("operation_time", _operationTime)
    return nil
}

// OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetOperationTime() string {
    return r._operationTime
}