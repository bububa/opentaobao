package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店库存初始化前后端商品绑定 API请求
tmall.nr.inventory.initial

用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式
*/
type TmallNrInventoryInitialAPIRequest struct {
    model.Params
    // 请求入参
    _param0   *NrStoreInvItemInitialReqDTO
}

// 初始化TmallNrInventoryInitialAPIRequest对象
func NewTmallNrInventoryInitialRequest() *TmallNrInventoryInitialAPIRequest{
    return &TmallNrInventoryInitialAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrInventoryInitialAPIRequest) GetApiMethodName() string {
    return "tmall.nr.inventory.initial"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrInventoryInitialAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 请求入参
func (r *TmallNrInventoryInitialAPIRequest) SetParam0(_param0 *NrStoreInvItemInitialReqDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallNrInventoryInitialAPIRequest) GetParam0() *NrStoreInvItemInitialReqDTO {
    return r._param0
}
