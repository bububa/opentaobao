package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingPicturesRecognize 图片识别
// taobao.growth.reaching.pictures.recognize
//
// 图片识别
func TaobaoGrowthReachingPicturesRecognize(clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingPicturesRecognizeAPIRequest, session string) (*usergrowth.TaobaoGrowthReachingPicturesRecognizeAPIResponse, error) {
	var resp usergrowth.TaobaoGrowthReachingPicturesRecognizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
