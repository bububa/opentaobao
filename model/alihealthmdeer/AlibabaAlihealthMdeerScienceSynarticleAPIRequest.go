package alihealthmdeer

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMdeerScienceSynarticleAPIRequest 医知鹿文章同步【保存/更新】 API请求
// alibaba.alihealth.mdeer.science.synarticle
//
// 文章同步【保存/更新】
type AlibabaAlihealthMdeerScienceSynarticleAPIRequest struct {
	model.Params
	// 同步文章对象
	_synArticleInfo *SynArticleInfo
}

// NewAlibabaAlihealthMdeerScienceSynarticleRequest 初始化AlibabaAlihealthMdeerScienceSynarticleAPIRequest对象
func NewAlibabaAlihealthMdeerScienceSynarticleRequest() *AlibabaAlihealthMdeerScienceSynarticleAPIRequest {
	return &AlibabaAlihealthMdeerScienceSynarticleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMdeerScienceSynarticleAPIRequest) Reset() {
	r._synArticleInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceSynarticleAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.science.synarticle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceSynarticleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMdeerScienceSynarticleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSynArticleInfo is SynArticleInfo Setter
// 同步文章对象
func (r *AlibabaAlihealthMdeerScienceSynarticleAPIRequest) SetSynArticleInfo(_synArticleInfo *SynArticleInfo) error {
	r._synArticleInfo = _synArticleInfo
	r.Set("syn_article_info", _synArticleInfo)
	return nil
}

// GetSynArticleInfo SynArticleInfo Getter
func (r AlibabaAlihealthMdeerScienceSynarticleAPIRequest) GetSynArticleInfo() *SynArticleInfo {
	return r._synArticleInfo
}

var poolAlibabaAlihealthMdeerScienceSynarticleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMdeerScienceSynarticleRequest()
	},
}

// GetAlibabaAlihealthMdeerScienceSynarticleRequest 从 sync.Pool 获取 AlibabaAlihealthMdeerScienceSynarticleAPIRequest
func GetAlibabaAlihealthMdeerScienceSynarticleAPIRequest() *AlibabaAlihealthMdeerScienceSynarticleAPIRequest {
	return poolAlibabaAlihealthMdeerScienceSynarticleAPIRequest.Get().(*AlibabaAlihealthMdeerScienceSynarticleAPIRequest)
}

// ReleaseAlibabaAlihealthMdeerScienceSynarticleAPIRequest 将 AlibabaAlihealthMdeerScienceSynarticleAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMdeerScienceSynarticleAPIRequest(v *AlibabaAlihealthMdeerScienceSynarticleAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMdeerScienceSynarticleAPIRequest.Put(v)
}
