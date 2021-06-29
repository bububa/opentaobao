package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建资金池 APIRequest
tmall.fans.cashpool.create

商家创建资金池接口
*/
type TmallFansCashpoolCreateRequest struct {
    model.Params

    // 创建资奖池输入对象
    createCashPoolParamDo   *CreateCashPoolParamDo 

}

func NewTmallFansCashpoolCreateRequest() *TmallFansCashpoolCreateRequest{
    return &TmallFansCashpoolCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallFansCashpoolCreateRequest) GetApiMethodName() string {
    return "tmall.fans.cashpool.create"
}

func (r TmallFansCashpoolCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallFansCashpoolCreateRequest) SetCreateCashPoolParamDo(createCashPoolParamDo *CreateCashPoolParamDo) error {
    r.createCashPoolParamDo = createCashPoolParamDo
    r.Set("create_cash_pool_param_do", createCashPoolParamDo)
    return nil
}

func (r TmallFansCashpoolCreateRequest) GetCreateCashPoolParamDo() *CreateCashPoolParamDo {
    return r.createCashPoolParamDo
}

