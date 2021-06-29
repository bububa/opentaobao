package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店业务同步库存 API请求
tmall.nr.inventory.update

用于商家每日同步更新门店库存
*/
type TmallNrInventoryUpdateRequest struct {
    model.Params
    // 入参
    _param0   *NrInventoryUpdateReqDto
}

// 初始化TmallNrInventoryUpdateRequest对象
func NewTmallNrInventoryUpdateRequest() *TmallNrInventoryUpdateRequest{
    return &TmallNrInventoryUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrInventoryUpdateRequest) GetApiMethodName() string {
    return "tmall.nr.inventory.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrInventoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TmallNrInventoryUpdateRequest) SetParam0(_param0 *NrInventoryUpdateReqDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallNrInventoryUpdateRequest) GetParam0() *NrInventoryUpdateReqDto {
    return r._param0
}
