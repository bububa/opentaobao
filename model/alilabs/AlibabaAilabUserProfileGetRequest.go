package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户信息 API请求
alibaba.ailab.user.profile.get

提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用
*/
type AlibabaAilabUserProfileGetRequest struct {
    model.Params
    // open uid
    openUid   string
    // client id
    clientId   string
}

// 初始化AlibabaAilabUserProfileGetRequest对象
func NewAlibabaAilabUserProfileGetRequest() *AlibabaAilabUserProfileGetRequest{
    return &AlibabaAilabUserProfileGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserProfileGetRequest) GetApiMethodName() string {
    return "alibaba.ailab.user.profile.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserProfileGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenUid Setter
// open uid
func (r *AlibabaAilabUserProfileGetRequest) SetOpenUid(openUid string) error {
    r.openUid = openUid
    r.Set("open_uid", openUid)
    return nil
}

// OpenUid Getter
func (r AlibabaAilabUserProfileGetRequest) GetOpenUid() string {
    return r.openUid
}
// ClientId Setter
// client id
func (r *AlibabaAilabUserProfileGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabUserProfileGetRequest) GetClientId() string {
    return r.clientId
}
