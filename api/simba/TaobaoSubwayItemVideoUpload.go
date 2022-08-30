package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayItemVideoUpload 创意视频上传
// taobao.subway.item.video.upload
//
// 为用户提供视频上传的功能
func TaobaoSubwayItemVideoUpload(clt *core.SDKClient, req *simba.TaobaoSubwayItemVideoUploadAPIRequest, session string) (*simba.TaobaoSubwayItemVideoUploadAPIResponse, error) {
	var resp simba.TaobaoSubwayItemVideoUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
