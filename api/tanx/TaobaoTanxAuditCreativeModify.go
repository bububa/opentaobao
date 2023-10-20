package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxauditcreativemodify 创意修改接口
// taobao.tanx.audit.creative.modify
//
// 创意修改接口
func Taobaotanxauditcreativemodify(clt *core.SDKClient, req *tanx.TaobaotanxauditcreativemodifyAPIRequest, session string) (*tanx.TaobaotanxauditcreativemodifyAPIResponse, error) {
	var resp tanx.TaobaotanxauditcreativemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
