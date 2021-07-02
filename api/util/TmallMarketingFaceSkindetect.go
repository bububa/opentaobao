package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TmallMarketingFaceSkindetect 肌肤检测
// tmall.marketing.face.skindetect
//
// 提供人脸肌肤属性报告
func TmallMarketingFaceSkindetect(clt *core.SDKClient, req *util.TmallMarketingFaceSkindetectAPIRequest, session string) (*util.TmallMarketingFaceSkindetectAPIResponse, error) {
	var resp util.TmallMarketingFaceSkindetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
