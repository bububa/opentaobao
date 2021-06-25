package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘口令配置数据 APIRequest
alibaba.baichuan.taopassword.config

百川淘口令规则配置接口
*/
type AlibabaBaichuanTaopasswordConfigRequest struct {
    model.Params

}

func NewAlibabaBaichuanTaopasswordConfigRequest() *AlibabaBaichuanTaopasswordConfigRequest{
    return &AlibabaBaichuanTaopasswordConfigRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBaichuanTaopasswordConfigRequest) GetApiMethodName() string {
    return "alibaba.baichuan.taopassword.config"
}

func (r AlibabaBaichuanTaopasswordConfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


