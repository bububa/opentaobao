package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuanctgvideoupload 提供优酷的短视频入淘API
// alibaba.baichuan.ctg.video.upload
//
// 提供优酷的短视频入淘API
func Alibababaichuanctgvideoupload(clt *core.SDKClient, req *baichuan.AlibababaichuanctgvideouploadAPIRequest, session string) (*baichuan.AlibababaichuanctgvideouploadAPIResponse, error) {
	var resp baichuan.AlibababaichuanctgvideouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
