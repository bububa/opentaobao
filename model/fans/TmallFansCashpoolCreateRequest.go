package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建资金池 API请求
tmall.fans.cashpool.create

商家创建资金池接口
*/
type TmallFansCashpoolCreateRequest struct {
    model.Params
    // 创建资奖池输入对象
    _createCashPoolParamDo   *CreateCashPoolParamDO
}

// 初始化TmallFansCashpoolCreateRequest对象
func NewTmallFansCashpoolCreateRequest() *TmallFansCashpoolCreateRequest{
    return &TmallFansCashpoolCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFansCashpoolCreateRequest) GetApiMethodName() string {
    return "tmall.fans.cashpool.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallFansCashpoolCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateCashPoolParamDo Setter
// 创建资奖池输入对象
func (r *TmallFansCashpoolCreateRequest) SetCreateCashPoolParamDo(_createCashPoolParamDo *CreateCashPoolParamDO) error {
    r._createCashPoolParamDo = _createCashPoolParamDo
    r.Set("create_cash_pool_param_do", _createCashPoolParamDo)
    return nil
}

// CreateCashPoolParamDo Getter
func (r TmallFansCashpoolCreateRequest) GetCreateCashPoolParamDo() *CreateCashPoolParamDO {
    return r._createCashPoolParamDo
}
