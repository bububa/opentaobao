package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘口令配置数据 API请求
alibaba.baichuan.taopassword.config

百川淘口令规则配置接口
*/
type AlibabaBaichuanTaopasswordConfigRequest struct {
    model.Params
}

// 初始化AlibabaBaichuanTaopasswordConfigRequest对象
func NewAlibabaBaichuanTaopasswordConfigRequest() *AlibabaBaichuanTaopasswordConfigRequest{
    return &AlibabaBaichuanTaopasswordConfigRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordConfigRequest) GetApiMethodName() string {
    return "alibaba.baichuan.taopassword.config"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordConfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
