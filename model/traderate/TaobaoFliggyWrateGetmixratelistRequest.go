package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪通用评价接口 APIRequest
taobao.fliggy.wrate.getmixratelist

飞猪评价通用接口
*/
type TaobaoFliggyWrateGetmixratelistRequest struct {
    model.Params

    // 评论列表查询参数
    paramTopGetMixRateListParam   *TopGetMixRateListParam 

}

func NewTaobaoFliggyWrateGetmixratelistRequest() *TaobaoFliggyWrateGetmixratelistRequest{
    return &TaobaoFliggyWrateGetmixratelistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFliggyWrateGetmixratelistRequest) GetApiMethodName() string {
    return "taobao.fliggy.wrate.getmixratelist"
}

func (r TaobaoFliggyWrateGetmixratelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFliggyWrateGetmixratelistRequest) SetParamTopGetMixRateListParam(paramTopGetMixRateListParam *TopGetMixRateListParam) error {
    r.paramTopGetMixRateListParam = paramTopGetMixRateListParam
    r.Set("param_top_get_mix_rate_list_param", paramTopGetMixRateListParam)
    return nil
}

func (r TaobaoFliggyWrateGetmixratelistRequest) GetParamTopGetMixRateListParam() *TopGetMixRateListParam {
    return r.paramTopGetMixRateListParam
}

