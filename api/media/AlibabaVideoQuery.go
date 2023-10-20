package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaVideoQuery 查询视频信息
// alibaba.video.query
//
// 查询视频信息
func AlibabaVideoQuery(clt *core.SDKClient, req *media.AlibabaVideoQueryAPIRequest, resp *media.AlibabaVideoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
