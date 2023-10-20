package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentmediaupload 闲鱼多媒体上传接口
// alibaba.idle.rent.media.upload
//
// 上传多媒体信息，包括图片、视频（暂不支持）
func Alibabaidlerentmediaupload(clt *core.SDKClient, req *idle.AlibabaidlerentmediauploadAPIRequest, session string) (*idle.AlibabaidlerentmediauploadAPIResponse, error) {
	var resp idle.AlibabaidlerentmediauploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
