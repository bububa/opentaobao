package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取海王用户权限包 API请求
alibaba.seaking.servicepack

获取海王用户权限包
*/
type AlibabaSeakingServicepackRequest struct {
    model.Params
    // 验证类型
    identifyType   string
    // 验证类型下的唯一id
    identifier   string
}

// 初始化AlibabaSeakingServicepackRequest对象
func NewAlibabaSeakingServicepackRequest() *AlibabaSeakingServicepackRequest{
    return &AlibabaSeakingServicepackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingServicepackRequest) GetApiMethodName() string {
    return "alibaba.seaking.servicepack"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingServicepackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdentifyType Setter
// 验证类型
func (r *AlibabaSeakingServicepackRequest) SetIdentifyType(identifyType string) error {
    r.identifyType = identifyType
    r.Set("identify_type", identifyType)
    return nil
}

// IdentifyType Getter
func (r AlibabaSeakingServicepackRequest) GetIdentifyType() string {
    return r.identifyType
}
// Identifier Setter
// 验证类型下的唯一id
func (r *AlibabaSeakingServicepackRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingServicepackRequest) GetIdentifier() string {
    return r.identifier
}
