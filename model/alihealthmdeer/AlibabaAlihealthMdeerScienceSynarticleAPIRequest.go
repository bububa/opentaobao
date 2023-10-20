package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmdeersciencesynarticleAPIRequest 医知鹿文章同步【保存/更新】 API请求
// alibaba.alihealth.mdeer.science.synarticle
//
// 文章同步【保存/更新】
type AlibabaalihealthmdeersciencesynarticleAPIRequest struct {
	model.Params
	// 同步文章对象
	_synArticleInfo *SynArticleInfo
}

// NewAlibabaalihealthmdeersciencesynarticleRequest 初始化AlibabaalihealthmdeersciencesynarticleAPIRequest对象
func NewAlibabaalihealthmdeersciencesynarticleRequest() *AlibabaalihealthmdeersciencesynarticleAPIRequest {
	return &AlibabaalihealthmdeersciencesynarticleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmdeersciencesynarticleAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.science.synarticle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmdeersciencesynarticleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmdeersciencesynarticleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSynArticleInfo is SynArticleInfo Setter
// 同步文章对象
func (r *AlibabaalihealthmdeersciencesynarticleAPIRequest) SetSynArticleInfo(_synArticleInfo *SynArticleInfo) error {
	r._synArticleInfo = _synArticleInfo
	r.Set("syn_article_info", _synArticleInfo)
	return nil
}

// GetSynArticleInfo SynArticleInfo Getter
func (r AlibabaalihealthmdeersciencesynarticleAPIRequest) GetSynArticleInfo() *SynArticleInfo {
	return r._synArticleInfo
}
