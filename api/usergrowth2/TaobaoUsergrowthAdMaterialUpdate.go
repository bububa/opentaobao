package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// Taobaousergrowthadmaterialupdate 素材更新
// taobao.usergrowth.ad.material.update
//
// 素材更新
func Taobaousergrowthadmaterialupdate(clt *core.SDKClient, req *usergrowth2.TaobaousergrowthadmaterialupdateAPIRequest, session string) (*usergrowth2.TaobaousergrowthadmaterialupdateAPIResponse, error) {
	var resp usergrowth2.TaobaousergrowthadmaterialupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
