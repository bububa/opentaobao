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
type TmallNrInventoryInitialRequest struct {
    model.Params
    // 请求入参
    param0   *NrStoreInvItemInitialReqDto
}

// 初始化TmallNrInventoryInitialRequest对象
func NewTmallNrInventoryInitialRequest() *TmallNrInventoryInitialRequest{
    return &TmallNrInventoryInitialRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrInventoryInitialRequest) GetApiMethodName() string {
    return "tmall.nr.inventory.initial"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrInventoryInitialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 请求入参
func (r *TmallNrInventoryInitialRequest) SetParam0(param0 *NrStoreInvItemInitialReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TmallNrInventoryInitialRequest) GetParam0() *NrStoreInvItemInitialReqDto {
    return r.param0
}
