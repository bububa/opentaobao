package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMdeerScienceSynVideoAPIRequest
视频同步【保存/更新】 API请求
alibaba.alihealth.mdeer.science.synVideo

视频同步【保存/更新】 */
type AlibabaAlihealthMdeerScienceSynVideoAPIRequest struct {
	model.Params
	// 视频信息实体
	_synVideoInfo *SynVideoInfo
}

// New
