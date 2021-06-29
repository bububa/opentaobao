package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店库存初始化前后端商品绑定 APIRequest
tmall.nr.inventory.initial

用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式
*/
type TmallNrInventoryInitialRequest struct {
    model.Params

    // 请求入参
    param0   *NrStoreInvItemInitialReqDto 

}

func NewTmallNrInventoryInitialRequest() *TmallNrInventoryInitialRequest{
    return &TmallNrInventoryInitialRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrInventoryInitialRequest) GetApiMethodName() string {
    return "tmall.nr.inventory.initial"
}

func (r TmallNrInventoryInitialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrInventoryInitialRequest) SetParam0(param0 *NrStoreInvItemInitialReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallNrInventoryInitialRequest) GetParam0() *NrStoreInvItemInitialReqDto {
    return r.param0
}

