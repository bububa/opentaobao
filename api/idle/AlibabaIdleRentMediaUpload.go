package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentMediaUpload 闲鱼多媒体上传接口
// alibaba.idle.rent.media.upload
//
// 上传多媒体信息，包括图片、视频（暂不支持）
func AlibabaIdleRentMediaUpload(clt *core.SDKClient, req *idle.AlibabaIdleRentMediaUploadAPIRequest, resp *idle.AlibabaIdleRentMediaUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
