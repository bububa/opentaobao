package alihealthmdeer

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// AlibabaAlihealthMdeerScienceSynVideo 视频同步【保存/更新】
// alibaba.alihealth.mdeer.science.synVideo
//
// 视频同步【保存/更新】
func AlibabaAlihealthMdeerScienceSynVideo(clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerScienceSynVideoAPIRequest, resp *alihealthmdeer.AlibabaAlihealthMdeerScienceSynVideoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
