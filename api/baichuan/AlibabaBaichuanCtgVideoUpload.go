package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanCtgVideoUpload 提供优酷的短视频入淘API
// alibaba.baichuan.ctg.video.upload
//
// 提供优酷的短视频入淘API
func AlibabaBaichuanCtgVideoUpload(clt *core.SDKClient, req *baichuan.AlibabaBaichuanCtgVideoUploadAPIRequest, session string) (*baichuan.AlibabaBaichuanCtgVideoUploadAPIResponse, error) {
	var resp baichuan.AlibabaBaichuanCtgVideoUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
