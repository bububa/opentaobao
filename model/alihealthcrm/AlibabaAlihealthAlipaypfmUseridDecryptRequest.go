package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户标识解密接口 API请求
alibaba.alihealth.alipaypfm.userid.decrypt

用户唯一表示加密传输，调用方解密
*/
type AlibabaAlihealthAlipaypfmUseridDecryptRequest struct {
    model.Params
    // 小程序appid
    _appId   string
    // 加密后的userId
    _content   string
}

// 初始化AlibabaAlihealthAlipaypfmUseridDecryptRequest对象
func NewAlibabaAlihealthAlipaypfmUseridDecryptRequest() *AlibabaAlihealthAlipaypfmUseridDecryptRequest{
    return &AlibabaAlihealthAlipaypfmUseridDecryptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.userid.decrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 小程序appid
func (r *AlibabaAlihealthAlipaypfmUseridDecryptRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetAppId() string {
    return r._appId
}
// Content Setter
// 加密后的userId
func (r *AlibabaAlihealthAlipaypfmUseridDecryptRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaAlihealthAlipaypfmUseridDecryptRequest) GetContent() string {
    return r._content
}
