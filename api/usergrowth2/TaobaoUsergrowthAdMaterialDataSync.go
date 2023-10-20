package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// Taobaousergrowthadmaterialdatasync 素材投放效果数据回传
// taobao.usergrowth.ad.material.data.sync
//
// 创意维度广告效果数据回传
func Taobaousergrowthadmaterialdatasync(clt *core.SDKClient, req *usergrowth2.TaobaousergrowthadmaterialdatasyncAPIRequest, session string) (*usergrowth2.TaobaousergrowthadmaterialdatasyncAPIResponse, error) {
	var resp usergrowth2.TaobaousergrowthadmaterialdatasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
