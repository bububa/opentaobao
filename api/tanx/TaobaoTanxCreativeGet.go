package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxcreativeget 获取单个审核创意状态
// taobao.tanx.creative.get
//
// 获取单个审核创意状态
func Taobaotanxcreativeget(clt *core.SDKClient, req *tanx.TaobaotanxcreativegetAPIRequest, session string) (*tanx.TaobaotanxcreativegetAPIResponse, error) {
	var resp tanx.TaobaotanxcreativegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
