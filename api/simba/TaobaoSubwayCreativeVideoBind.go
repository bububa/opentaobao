package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCreativeVideoBind 绑定视频到创意上
// taobao.subway.creative.video.bind
//
// 将用户上传的视频绑定到指定的创意上
func TaobaoSubwayCreativeVideoBind(clt *core.SDKClient, req *simba.TaobaoSubwayCreativeVideoBindAPIRequest, session string) (*simba.TaobaoSubwayCreativeVideoBindAPIResponse, error) {
	var resp simba.TaobaoSubwayCreativeVideoBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
