package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomevideosync 视频草稿信息同步
// alibaba.alihouse.newhome.video.sync
//
// 接收视频信息记录
func Alibabaalihousenewhomevideosync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomevideosyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomevideosyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomevideosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
