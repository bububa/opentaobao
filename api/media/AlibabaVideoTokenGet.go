package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaVideoTokenGet 获取上传token
// alibaba.video.token.get
//
// 获取上传token
func AlibabaVideoTokenGet(clt *core.SDKClient, req *media.AlibabaVideoTokenGetAPIRequest, resp *media.AlibabaVideoTokenGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
