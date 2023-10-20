package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Alibabavideoquery 查询视频信息
// alibaba.video.query
//
// 查询视频信息
func Alibabavideoquery(clt *core.SDKClient, req *media.AlibabavideoqueryAPIRequest, session string) (*media.AlibabavideoqueryAPIResponse, error) {
	var resp media.AlibabavideoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
