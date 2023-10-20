package yunosad

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

// Yunosadauditcreativegetlist 批量获取创意审核状态
// yunos.ad.audit.creative.getlist
//
// 批量获取创意审核状态
func Yunosadauditcreativegetlist(clt *core.SDKClient, req *yunosad.YunosadauditcreativegetlistAPIRequest, session string) (*yunosad.YunosadauditcreativegetlistAPIResponse, error) {
	var resp yunosad.YunosadauditcreativegetlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
