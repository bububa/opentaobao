package yunosad

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

// Yunosadauditcreativeget 获取单个创意审核状态
// yunos.ad.audit.creative.get
//
// 获取单个创意审核状态
func Yunosadauditcreativeget(clt *core.SDKClient, req *yunosad.YunosadauditcreativegetAPIRequest, session string) (*yunosad.YunosadauditcreativegetAPIResponse, error) {
	var resp yunosad.YunosadauditcreativegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
