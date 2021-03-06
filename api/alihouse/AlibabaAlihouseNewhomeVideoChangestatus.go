package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeVideoChangestatus 视频草稿状态更新
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
func AlibabaAlihouseNewhomeVideoChangestatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeVideoChangestatusAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeVideoChangestatusAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeVideoChangestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
