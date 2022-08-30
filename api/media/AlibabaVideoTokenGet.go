package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaVideoTokenGet 获取上传token
// alibaba.video.token.get
//
// 获取上传token
func AlibabaVideoTokenGet(clt *core.SDKClient, req *media.AlibabaVideoTokenGetAPIRequest, session string) (*media.AlibabaVideoTokenGetAPIResponse, error) {
	var resp media.AlibabaVideoTokenGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
