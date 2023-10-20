package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxauditcreativeadd 创意预审新增接口
// taobao.tanx.audit.creative.add
//
// 创意预审新增接口
func Taobaotanxauditcreativeadd(clt *core.SDKClient, req *tanx.TaobaotanxauditcreativeaddAPIRequest, session string) (*tanx.TaobaotanxauditcreativeaddAPIResponse, error) {
	var resp tanx.TaobaotanxauditcreativeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
