package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// Taobaousergrowthadmaterialaudit 素材审核
// taobao.usergrowth.ad.material.audit
//
// 素材审核
func Taobaousergrowthadmaterialaudit(clt *core.SDKClient, req *usergrowth2.TaobaousergrowthadmaterialauditAPIRequest, session string) (*usergrowth2.TaobaousergrowthadmaterialauditAPIResponse, error) {
	var resp usergrowth2.TaobaousergrowthadmaterialauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
