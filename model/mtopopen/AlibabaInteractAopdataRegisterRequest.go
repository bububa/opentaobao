package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资源位数据推送接口 APIRequest
alibaba.interact.aopdata.register

提供给isv，查询以及推送浮层资源位的三方活动数据
*/
type AlibabaInteractAopdataRegisterRequest struct {
    model.Params

    // 入参
    paramTopIsvDecorateParam   *TopIsvDecorateParam 

}

func NewAlibabaInteractAopdataRegisterRequest() *AlibabaInteractAopdataRegisterRequest{
    return &AlibabaInteractAopdataRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractAopdataRegisterRequest) GetApiMethodName() string {
    return "alibaba.interact.aopdata.register"
}

func (r AlibabaInteractAopdataRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractAopdataRegisterRequest) SetParamTopIsvDecorateParam(paramTopIsvDecorateParam *TopIsvDecorateParam) error {
    r.paramTopIsvDecorateParam = paramTopIsvDecorateParam
    r.Set("param_top_isv_decorate_param", paramTopIsvDecorateParam)
    return nil
}

func (r AlibabaInteractAopdataRegisterRequest) GetParamTopIsvDecorateParam() *TopIsvDecorateParam {
    return r.paramTopIsvDecorateParam
}

