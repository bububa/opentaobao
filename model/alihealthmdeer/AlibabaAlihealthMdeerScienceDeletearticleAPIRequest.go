package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMdeerScienceDeletearticleAPIRequest
文章删除 API请求
alibaba.alihealth.mdeer.science.deletearticle

三方同步文章删除 */
type AlibabaAlihealthMdeerScienceDeletearticleAPIRequest struct {
	model.Params
	// 文章ID
	_articleId int64
}

// New
