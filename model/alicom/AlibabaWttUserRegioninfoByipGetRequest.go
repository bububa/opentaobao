package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ip获取省市信息 API请求
alibaba.wtt.user.regioninfo.byip.get

通过ip获取省市信息
*/
type AlibabaWttUserRegioninfoByipGetRequest struct {
    model.Params
    // ip地址
    ip   string
}

// 初始化AlibabaWttUserRegioninfoByipGetRequest对象
func NewAlibabaWttUserRegioninfoByipGetRequest() *AlibabaWttUserRegioninfoByipGetRequest{
    return &AlibabaWttUserRegioninfoByipGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWttUserRegioninfoByipGetRequest) GetApiMethodName() string {
    return "alibaba.wtt.user.regioninfo.byip.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWttUserRegioninfoByipGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ip Setter
// ip地址
func (r *AlibabaWttUserRegioninfoByipGetRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaWttUserRegioninfoByipGetRequest) GetIp() string {
    return r.ip
}
