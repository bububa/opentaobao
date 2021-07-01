package aliexpress

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步帖子 API请求
aliexpress.feed.post.publish

站外平台同步帖子至AE FEED域
*/
type AliexpressFeedPostPublishAPIRequest struct {
    model.Params
    // 站外导入内容请求参数
    _offsitePublishPostEntity   *OffsitePublishPostEntity
}

// 初始化AliexpressFeedPostPublishAPIRequest对象
func NewAliexpressFeedPostPublishRequest() *AliexpressFeedPostPublishAPIRequest{
    return &AliexpressFeedPostPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressFeedPostPublishAPIRequest) GetApiMethodName() string {
    return "aliexpress.feed.post.publish"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressFeedPostPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OffsitePublishPostEntity Setter
// 站外导入内容请求参数
func (r *AliexpressFeedPostPublishAPIRequest) SetOffsitePublishPostEntity(_offsitePublishPostEntity *OffsitePublishPostEntity) error {
    r._offsitePublishPostEntity = _offsitePublishPostEntity
    r.Set("offsite_publish_post_entity", _offsitePublishPostEntity)
    return nil
}

// OffsitePublishPostEntity Getter
func (r AliexpressFeedPostPublishAPIRequest) GetOffsitePublishPostEntity() *OffsitePublishPostEntity {
    return r._offsitePublishPostEntity
}
