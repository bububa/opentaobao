package usergrowth

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingPicturesRecognize 图片识别
// taobao.growth.reaching.pictures.recognize
//
// 图片识别
func TaobaoGrowthReachingPicturesRecognize(ctx context.Context, clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingPicturesRecognizeAPIRequest, resp *usergrowth.TaobaoGrowthReachingPicturesRecognizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
