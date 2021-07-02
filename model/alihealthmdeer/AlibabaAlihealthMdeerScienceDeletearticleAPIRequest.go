package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMdeerScienceDeletearticleAPIRequest 文章删除 API请求
// alibaba.alihealth.mdeer.science.deletearticle
//
// 三方同步文章删除
type AlibabaAlihealthMdeerScienceDeletearticleAPIRequest struct {
	model.Params
	// 文章ID
	_articleId int64
}

// NewAlibabaAlihealthMdeerScienceDeletearticleRequest 初始化AlibabaAlihealthMdeerScienceDeletearticleAPIRequest对象
func NewAlibabaAlihealthMdeerScienceDeletearticleRequest() *AlibabaAlihealthMdeerScienceDeletearticleAPIRequest {
	return &AlibabaAlihealthMdeerScienceDeletearticleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.science.deletearticle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetArticleId is ArticleId Setter
// 文章ID
func (r *AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) SetArticleId(_articleId int64) error {
	r._articleId = _articleId
	r.Set("article_id", _articleId)
	return nil
}

// GetArticleId ArticleId Getter
func (r AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) GetArticleId() int64 {
	return r._articleId
}
