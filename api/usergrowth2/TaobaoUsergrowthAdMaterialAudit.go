package usergrowth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialAudit 素材审核
// taobao.usergrowth.ad.material.audit
//
// 素材审核
func TaobaoUsergrowthAdMaterialAudit(ctx context.Context, clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialAuditAPIRequest, resp *usergrowth2.TaobaoUsergrowthAdMaterialAuditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
