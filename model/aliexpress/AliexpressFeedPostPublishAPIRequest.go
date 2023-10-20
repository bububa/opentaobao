package aliexpress

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressfeedpostpublishAPIRequest 同步帖子 API请求
// aliexpress.feed.post.publish
//
// 站外平台同步帖子至AE FEED域
type AliexpressfeedpostpublishAPIRequest struct {
	model.Params
	// 站外导入内容请求参数
	_offsitePublishPostEntity *OffsitePublishPostEntity
}

// NewAliexpressfeedpostpublishRequest 初始化AliexpressfeedpostpublishAPIRequest对象
func NewAliexpressfeedpostpublishRequest() *AliexpressfeedpostpublishAPIRequest {
	return &AliexpressfeedpostpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressfeedpostpublishAPIRequest) GetApiMethodName() string {
	return "aliexpress.feed.post.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressfeedpostpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressfeedpostpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOffsitePublishPostEntity is OffsitePublishPostEntity Setter
// 站外导入内容请求参数
func (r *AliexpressfeedpostpublishAPIRequest) SetOffsitePublishPostEntity(_offsitePublishPostEntity *OffsitePublishPostEntity) error {
	r._offsitePublishPostEntity = _offsitePublishPostEntity
	r.Set("offsite_publish_post_entity", _offsitePublishPostEntity)
	return nil
}

// GetOffsitePublishPostEntity OffsitePublishPostEntity Getter
func (r AliexpressfeedpostpublishAPIRequest) GetOffsitePublishPostEntity() *OffsitePublishPostEntity {
	return r._offsitePublishPostEntity
}
