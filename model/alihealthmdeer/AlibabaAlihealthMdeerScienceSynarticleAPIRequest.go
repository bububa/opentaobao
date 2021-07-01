package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMdeerScienceSynarticleAPIRequest
医知鹿文章同步【保存/更新】 API请求
alibaba.alihealth.mdeer.science.synarticle

文章同步【保存/更新】 */
type AlibabaAlihealthMdeerScienceSynarticleAPIRequest struct {
	model.Params
	// 同步文章对象
	_synArticleInfo *SynArticleInfo
}

// New
