package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

/* AlibabaScsImageMatte
阿里妈妈智能创意平台在线抠图
alibaba.scs.image.matte

该API对外输出一个在线抠图(Deep Image Matting)接口，合作方可以通过该接口利用深度学习抠图算法从图片中抠出目标对象(比如商品或者人物轮廓) */
func AlibabaScsImageMatte(clt *core.SDKClient, req *scs.AlibabaScsImageMatteAPIRequest, session string) (*scs.AlibabaScsImageMatteAPIResponse, error) {
	var resp scs.AlibabaScsImageMatteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
