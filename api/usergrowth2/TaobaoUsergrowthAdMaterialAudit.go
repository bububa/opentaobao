package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialAudit 素材审核
// taobao.usergrowth.ad.material.audit
//
// 素材审核
func TaobaoUsergrowthAdMaterialAudit(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialAuditAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthAdMaterialAuditAPIResponse, error) {
	var resp usergrowth2.TaobaoUsergrowthAdMaterialAuditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
