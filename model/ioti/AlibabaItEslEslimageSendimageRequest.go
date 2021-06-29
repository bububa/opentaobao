package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下发厂测初始化图片 APIRequest
alibaba.it.esl.eslimage.sendimage

工厂对生产出的电子价签进行全流程功能测试，能将出场图片通过ESL系统初始化到电子价签中
*/
type AlibabaItEslEslimageSendimageRequest struct {
    model.Params

    // 价签地址
    mac   string 

}

func NewAlibabaItEslEslimageSendimageRequest() *AlibabaItEslEslimageSendimageRequest{
    return &AlibabaItEslEslimageSendimageRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItEslEslimageSendimageRequest) GetApiMethodName() string {
    return "alibaba.it.esl.eslimage.sendimage"
}

func (r AlibabaItEslEslimageSendimageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItEslEslimageSendimageRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

func (r AlibabaItEslEslimageSendimageRequest) GetMac() string {
    return r.mac
}

