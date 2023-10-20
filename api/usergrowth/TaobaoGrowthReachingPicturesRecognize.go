package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaogrowthreachingpicturesrecognize 图片识别
// taobao.growth.reaching.pictures.recognize
//
// 图片识别
func Taobaogrowthreachingpicturesrecognize(clt *core.SDKClient, req *usergrowth.TaobaogrowthreachingpicturesrecognizeAPIRequest, session string) (*usergrowth.TaobaogrowthreachingpicturesrecognizeAPIResponse, error) {
	var resp usergrowth.TaobaogrowthreachingpicturesrecognizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
