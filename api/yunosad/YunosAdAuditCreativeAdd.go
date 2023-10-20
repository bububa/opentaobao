package yunosad

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

// Yunosadauditcreativeadd 单个创意预审接口
// yunos.ad.audit.creative.add
//
// YunOS广告业务第三方DSP单个创意预审接口
func Yunosadauditcreativeadd(clt *core.SDKClient, req *yunosad.YunosadauditcreativeaddAPIRequest, session string) (*yunosad.YunosadauditcreativeaddAPIResponse, error) {
	var resp yunosad.YunosadauditcreativeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
