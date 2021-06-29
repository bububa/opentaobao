package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调整库存 APIRequest
alibaba.mos.goods.adjust

库存调整接口
*/
type AlibabaMosGoodsAdjustRequest struct {
    model.Params

    // 库存调整入参
    paramIsvStockAdjustRequestDTO   *IsvStockAdjustRequestDto 

}

func NewAlibabaMosGoodsAdjustRequest() *AlibabaMosGoodsAdjustRequest{
    return &AlibabaMosGoodsAdjustRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosGoodsAdjustRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.adjust"
}

func (r AlibabaMosGoodsAdjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosGoodsAdjustRequest) SetParamIsvStockAdjustRequestDTO(paramIsvStockAdjustRequestDTO *IsvStockAdjustRequestDto) error {
    r.paramIsvStockAdjustRequestDTO = paramIsvStockAdjustRequestDTO
    r.Set("param_isv_stock_adjust_request_d_t_o", paramIsvStockAdjustRequestDTO)
    return nil
}

func (r AlibabaMosGoodsAdjustRequest) GetParamIsvStockAdjustRequestDTO() *IsvStockAdjustRequestDto {
    return r.paramIsvStockAdjustRequestDTO
}

