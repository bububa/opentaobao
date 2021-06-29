package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店业务同步库存 APIRequest
tmall.nr.inventory.update

用于商家每日同步更新门店库存
*/
type TmallNrInventoryUpdateRequest struct {
    model.Params

    // 入参
    param0   *NrInventoryUpdateReqDto 

}

func NewTmallNrInventoryUpdateRequest() *TmallNrInventoryUpdateRequest{
    return &TmallNrInventoryUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrInventoryUpdateRequest) GetApiMethodName() string {
    return "tmall.nr.inventory.update"
}

func (r TmallNrInventoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrInventoryUpdateRequest) SetParam0(param0 *NrInventoryUpdateReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallNrInventoryUpdateRequest) GetParam0() *NrInventoryUpdateReqDto {
    return r.param0
}

