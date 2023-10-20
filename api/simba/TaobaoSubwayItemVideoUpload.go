package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayitemvideoupload 创意视频上传
// taobao.subway.item.video.upload
//
// 为用户提供视频上传的功能
func Taobaosubwayitemvideoupload(clt *core.SDKClient, req *simba.TaobaosubwayitemvideouploadAPIRequest, session string) (*simba.TaobaosubwayitemvideouploadAPIResponse, error) {
	var resp simba.TaobaosubwayitemvideouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
