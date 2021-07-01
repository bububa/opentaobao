package yunosad

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

/* YunosAdAuditCreativeGet
获取单个创意审核状态
yunos.ad.audit.creative.get

获取单个创意审核状态 */
func YunosAdAuditCreativeGet(clt *core.SDKClient, req *yunosad.YunosAdAuditCreativeGetAPIRequest, session string) (*yunosad.YunosAdAuditCreativeGetAPIResponse, error) {
	var resp yunosad.YunosAdAuditCreativeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
