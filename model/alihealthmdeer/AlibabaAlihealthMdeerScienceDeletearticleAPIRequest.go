package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmdeersciencedeletearticleAPIRequest 文章删除 API请求
// alibaba.alihealth.mdeer.science.deletearticle
//
// 三方同步文章删除
type AlibabaalihealthmdeersciencedeletearticleAPIRequest struct {
	model.Params
	// 文章ID
	_articleId int64
}

// NewAlibabaalihealthmdeersciencedeletearticleRequest 初始化AlibabaalihealthmdeersciencedeletearticleAPIRequest对象
func NewAlibabaalihealthmdeersciencedeletearticleRequest() *AlibabaalihealthmdeersciencedeletearticleAPIRequest {
	return &AlibabaalihealthmdeersciencedeletearticleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmdeersciencedeletearticleAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.science.deletearticle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmdeersciencedeletearticleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmdeersciencedeletearticleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetArticleId is ArticleId Setter
// 文章ID
func (r *AlibabaalihealthmdeersciencedeletearticleAPIRequest) SetArticleId(_articleId int64) error {
	r._articleId = _articleId
	r.Set("article_id", _articleId)
	return nil
}

// GetArticleId ArticleId Getter
func (r AlibabaalihealthmdeersciencedeletearticleAPIRequest) GetArticleId() int64 {
	return r._articleId
}
