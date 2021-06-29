package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资源位数据推送接口 API请求
alibaba.interact.aopdata.register

提供给isv，查询以及推送浮层资源位的三方活动数据
*/
type AlibabaInteractAopdataRegisterRequest struct {
    model.Params
    // 入参
    _paramTopIsvDecorateParam   *TopIsvDecorateParam
}

// 初始化AlibabaInteractAopdataRegisterRequest对象
func NewAlibabaInteractAopdataRegisterRequest() *AlibabaInteractAopdataRegisterRequest{
    return &AlibabaInteractAopdataRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractAopdataRegisterRequest) GetApiMethodName() string {
    return "alibaba.interact.aopdata.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractAopdataRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopIsvDecorateParam Setter
// 入参
func (r *AlibabaInteractAopdataRegisterRequest) SetParamTopIsvDecorateParam(_paramTopIsvDecorateParam *TopIsvDecorateParam) error {
    r._paramTopIsvDecorateParam = _paramTopIsvDecorateParam
    r.Set("param_top_isv_decorate_param", _paramTopIsvDecorateParam)
    return nil
}

// ParamTopIsvDecorateParam Getter
func (r AlibabaInteractAopdataRegisterRequest) GetParamTopIsvDecorateParam() *TopIsvDecorateParam {
    return r._paramTopIsvDecorateParam
}
