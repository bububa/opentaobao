package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪通用评价接口 API请求
taobao.fliggy.wrate.getmixratelist

飞猪评价通用接口
*/
type TaobaoFliggyWrateGetmixratelistAPIRequest struct {
    model.Params
    // 评论列表查询参数
    _paramTopGetMixRateListParam   *TopGetMixRateListParam
}

// 初始化TaobaoFliggyWrateGetmixratelistAPIRequest对象
func NewTaobaoFliggyWrateGetmixratelistRequest() *TaobaoFliggyWrateGetmixratelistAPIRequest{
    return &TaobaoFliggyWrateGetmixratelistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFliggyWrateGetmixratelistAPIRequest) GetApiMethodName() string {
    return "taobao.fliggy.wrate.getmixratelist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFliggyWrateGetmixratelistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopGetMixRateListParam Setter
// 评论列表查询参数
func (r *TaobaoFliggyWrateGetmixratelistAPIRequest) SetParamTopGetMixRateListParam(_paramTopGetMixRateListParam *TopGetMixRateListParam) error {
    r._paramTopGetMixRateListParam = _paramTopGetMixRateListParam
    r.Set("param_top_get_mix_rate_list_param", _paramTopGetMixRateListParam)
    return nil
}

// ParamTopGetMixRateListParam Getter
func (r TaobaoFliggyWrateGetmixratelistAPIRequest) GetParamTopGetMixRateListParam() *TopGetMixRateListParam {
    return r._paramTopGetMixRateListParam
}
