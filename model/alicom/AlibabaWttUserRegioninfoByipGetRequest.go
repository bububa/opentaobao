package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ip获取省市信息 APIRequest
alibaba.wtt.user.regioninfo.byip.get

通过ip获取省市信息
*/
type AlibabaWttUserRegioninfoByipGetRequest struct {
    model.Params

    // ip地址
    ip   string 

}

func NewAlibabaWttUserRegioninfoByipGetRequest() *AlibabaWttUserRegioninfoByipGetRequest{
    return &AlibabaWttUserRegioninfoByipGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWttUserRegioninfoByipGetRequest) GetApiMethodName() string {
    return "alibaba.wtt.user.regioninfo.byip.get"
}

func (r AlibabaWttUserRegioninfoByipGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWttUserRegioninfoByipGetRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r AlibabaWttUserRegioninfoByipGetRequest) GetIp() string {
    return r.ip
}

