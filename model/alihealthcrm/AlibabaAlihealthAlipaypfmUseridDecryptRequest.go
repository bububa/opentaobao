package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户标识解密接口 APIRequest
alibaba.alihealth.alipaypfm.userid.decrypt

用户唯一表示加密传输，调用方解密
*/
type AlibabaAlihealthAlipaypfmUseridDecryptRequest struct {
    model.Params

    // 小程序appid
    appId   string 

    // 加密后的userId
    content   string 

}

func NewAlibabaAlihealthAlipaypfmUseridDecryptRequest() *AlibabaAlihealthAlipaypfmUseridDecryptRequest{
    return &AlibabaAlihealthAlipaypfmUseridDecryptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.userid.decrypt"
}

func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAlipaypfmUseridDecryptRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaAlihealthAlipaypfmUseridDecryptRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetContent() string {
    return r.content
}

