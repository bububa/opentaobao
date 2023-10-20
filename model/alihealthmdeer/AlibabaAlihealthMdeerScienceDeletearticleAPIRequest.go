package alihealthmdeer

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) Reset() {
	r._articleId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.science.deletearticle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthMdeerScienceDeletearticleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMdeerScienceDeletearticleRequest()
	},
}

// GetAlibabaAlihealthMdeerScienceDeletearticleRequest 从 sync.Pool 获取 AlibabaAlihealthMdeerScienceDeletearticleAPIRequest
func GetAlibabaAlihealthMdeerScienceDeletearticleAPIRequest() *AlibabaAlihealthMdeerScienceDeletearticleAPIRequest {
	return poolAlibabaAlihealthMdeerScienceDeletearticleAPIRequest.Get().(*AlibabaAlihealthMdeerScienceDeletearticleAPIRequest)
}

// ReleaseAlibabaAlihealthMdeerScienceDeletearticleAPIRequest 将 AlibabaAlihealthMdeerScienceDeletearticleAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMdeerScienceDeletearticleAPIRequest(v *AlibabaAlihealthMdeerScienceDeletearticleAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMdeerScienceDeletearticleAPIRequest.Put(v)
}
