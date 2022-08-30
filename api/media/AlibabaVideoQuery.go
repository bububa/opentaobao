package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaVideoQuery 查询视频信息
// alibaba.video.query
//
// 查询视频信息
func AlibabaVideoQuery(clt *core.SDKClient, req *media.AlibabaVideoQueryAPIRequest, session string) (*media.AlibabaVideoQueryAPIResponse, error) {
	var resp media.AlibabaVideoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
