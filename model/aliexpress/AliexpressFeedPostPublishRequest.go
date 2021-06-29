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
type AliexpressFeedPostPublishRequest struct {
    model.Params
    // 站外导入内容请求参数
    offsitePublishPostEntity   *OffsitePublishPostEntity
}

// 初始化AliexpressFeedPostPublishRequest对象
func NewAliexpressFeedPostPublishRequest() *AliexpressFeedPostPublishRequest{
    return &AliexpressFeedPostPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressFeedPostPublishRequest) GetApiMethodName() string {
    return "aliexpress.feed.post.publish"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressFeedPostPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OffsitePublishPostEntity Setter
// 站外导入内容请求参数
func (r *AliexpressFeedPostPublishRequest) SetOffsitePublishPostEntity(offsitePublishPostEntity *OffsitePublishPostEntity) error {
    r.offsitePublishPostEntity = offsitePublishPostEntity
    r.Set("offsite_publish_post_entity", offsitePublishPostEntity)
    return nil
}

// OffsitePublishPostEntity Getter
func (r AliexpressFeedPostPublishRequest) GetOffsitePublishPostEntity() *OffsitePublishPostEntity {
    return r.offsitePublishPostEntity
}
