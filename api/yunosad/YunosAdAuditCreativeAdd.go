package yunosad

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

// YunosAdAuditCreativeAdd 单个创意预审接口
// yunos.ad.audit.creative.add
//
// YunOS广告业务第三方DSP单个创意预审接口
func YunosAdAuditCreativeAdd(clt *core.SDKClient, req *yunosad.YunosAdAuditCreativeAddAPIRequest, session string) (*yunosad.YunosAdAuditCreativeAddAPIResponse, error) {
	var resp yunosad.YunosAdAuditCreativeAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
