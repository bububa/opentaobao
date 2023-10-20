package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// Taobaousergrowthadmaterialquery 素材审核结果查询
// taobao.usergrowth.ad.material.query
//
// 素材审核结果查询
func Taobaousergrowthadmaterialquery(clt *core.SDKClient, req *usergrowth2.TaobaousergrowthadmaterialqueryAPIRequest, session string) (*usergrowth2.TaobaousergrowthadmaterialqueryAPIResponse, error) {
	var resp usergrowth2.TaobaousergrowthadmaterialqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
