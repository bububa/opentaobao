package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Alibabavideotokenget 获取上传token
// alibaba.video.token.get
//
// 获取上传token
func Alibabavideotokenget(clt *core.SDKClient, req *media.AlibabavideotokengetAPIRequest, session string) (*media.AlibabavideotokengetAPIResponse, error) {
	var resp media.AlibabavideotokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
