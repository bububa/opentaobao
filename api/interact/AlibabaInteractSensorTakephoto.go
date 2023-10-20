package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensortakephoto takePhoto
// alibaba.interact.sensor.takephoto
//
// 客户端takePhoto
func Alibabainteractsensortakephoto(clt *core.SDKClient, req *interact.AlibabainteractsensortakephotoAPIRequest, session string) (*interact.AlibabainteractsensortakephotoAPIResponse, error) {
	var resp interact.AlibabainteractsensortakephotoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
