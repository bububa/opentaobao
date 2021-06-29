package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
厂测LED控制 API请求
alibaba.it.esl.sendled

针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
*/
type AlibabaItEslSendledRequest struct {
    model.Params
    // mac
    _macAp   string
    // 0、1、2、3：关蓝绿红
    _type   string
}

// 初始化AlibabaItEslSendledRequest对象
func NewAlibabaItEslSendledRequest() *AlibabaItEslSendledRequest{
    return &AlibabaItEslSendledRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItEslSendledRequest) GetApiMethodName() string {
    return "alibaba.it.esl.sendled"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItEslSendledRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MacAp Setter
// mac
func (r *AlibabaItEslSendledRequest) SetMacAp(_macAp string) error {
    r._macAp = _macAp
    r.Set("mac_ap", _macAp)
    return nil
}

// MacAp Getter
func (r AlibabaItEslSendledRequest) GetMacAp() string {
    return r._macAp
}
// Type Setter
// 0、1、2、3：关蓝绿红
func (r *AlibabaItEslSendledRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaItEslSendledRequest) GetType() string {
    return r._type
}
