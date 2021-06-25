package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序获取登录用户信息 APIRequest
alibaba.lsy.miniapp.user.get

零售云小程序，通过授权码获取登录的卖家账号信息
*/
type AlibabaLsyMiniappUserGetRequest struct {
    model.Params

    // 当前时间戳，毫秒
    timeStamp   string 

    // 获取用户信息的授权码，在小程序中获取
    code   string 

    // 请求参数签名，sha1(所有入参+appSecret，按字符串升序排列)
    signature   string 

    // 系统分配的小程序ID
    appId   string 

}

func NewAlibabaLsyMiniappUserGetRequest() *AlibabaLsyMiniappUserGetRequest{
    return &AlibabaLsyMiniappUserGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyMiniappUserGetRequest) GetApiMethodName() string {
    return "alibaba.lsy.miniapp.user.get"
}

func (r AlibabaLsyMiniappUserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyMiniappUserGetRequest) SetTimeStamp(timeStamp string) error {
    r.timeStamp = timeStamp
    r.Set("time_stamp", timeStamp)
    return nil
}

func (r AlibabaLsyMiniappUserGetRequest) GetTimeStamp() string {
    return r.timeStamp
}

func (r *AlibabaLsyMiniappUserGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaLsyMiniappUserGetRequest) GetCode() string {
    return r.code
}

func (r *AlibabaLsyMiniappUserGetRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

func (r AlibabaLsyMiniappUserGetRequest) GetSignature() string {
    return r.signature
}

func (r *AlibabaLsyMiniappUserGetRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaLsyMiniappUserGetRequest) GetAppId() string {
    return r.appId
}

