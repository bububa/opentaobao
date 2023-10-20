package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomevideochangestatus 视频草稿状态更新
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
func Alibabaalihousenewhomevideochangestatus(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomevideochangestatusAPIRequest, session string) (*alihouse.AlibabaalihousenewhomevideochangestatusAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomevideochangestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
