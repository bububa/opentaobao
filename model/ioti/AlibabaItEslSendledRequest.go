package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
厂测LED控制 APIRequest
alibaba.it.esl.sendled

针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
*/
type AlibabaItEslSendledRequest struct {
    model.Params

    // mac
    macAp   string 

    // 0、1、2、3：关蓝绿红
    type   string 

}

func NewAlibabaItEslSendledRequest() *AlibabaItEslSendledRequest{
    return &AlibabaItEslSendledRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItEslSendledRequest) GetApiMethodName() string {
    return "alibaba.it.esl.sendled"
}

func (r AlibabaItEslSendledRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItEslSendledRequest) SetMacAp(macAp string) error {
    r.macAp = macAp
    r.Set("mac_ap", macAp)
    return nil
}

func (r AlibabaItEslSendledRequest) GetMacAp() string {
    return r.macAp
}

func (r *AlibabaItEslSendledRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaItEslSendledRequest) GetType() string {
    return r.type
}

