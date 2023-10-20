package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxauditdepositcreativeadd dsp托管创意新增接口
// taobao.tanx.audit.depositcreative.add
//
// dsp托管创意新增接口
func Taobaotanxauditdepositcreativeadd(clt *core.SDKClient, req *tanx.TaobaotanxauditdepositcreativeaddAPIRequest, session string) (*tanx.TaobaotanxauditdepositcreativeaddAPIResponse, error) {
	var resp tanx.TaobaotanxauditdepositcreativeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
