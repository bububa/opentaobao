package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过外部订单ID查询所有资金单 API请求
taobao.alitrip.axin.trans.fund.query.by.order

阿信供销平台-通过外部订单ID查询所有资金单
*/
type TaobaoAlitripAxinTransFundQueryByOrderRequest struct {
    model.Params
    // 入参
    paramAxinFundListQueryDTO   *AxinFundListQueryDto
}

// 初始化TaobaoAlitripAxinTransFundQueryByOrderRequest对象
func NewTaobaoAlitripAxinTransFundQueryByOrderRequest() *TaobaoAlitripAxinTransFundQueryByOrderRequest{
    return &TaobaoAlitripAxinTransFundQueryByOrderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundQueryByOrderRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.fund.query.by.order"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundQueryByOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAxinFundListQueryDTO Setter
// 入参
func (r *TaobaoAlitripAxinTransFundQueryByOrderRequest) SetParamAxinFundListQueryDTO(paramAxinFundListQueryDTO *AxinFundListQueryDto) error {
    r.paramAxinFundListQueryDTO = paramAxinFundListQueryDTO
    r.Set("param_axin_fund_list_query_d_t_o", paramAxinFundListQueryDTO)
    return nil
}

// ParamAxinFundListQueryDTO Getter
func (r TaobaoAlitripAxinTransFundQueryByOrderRequest) GetParamAxinFundListQueryDTO() *AxinFundListQueryDto {
    return r.paramAxinFundListQueryDTO
}
