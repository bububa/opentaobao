package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgrouponlineitemsvonget 获取用户上架在线销售的全部宝贝
// taobao.simba.adgroup.onlineitemsvon.get
//
// 获取用户上架在线销售的全部宝贝
func Taobaosimbaadgrouponlineitemsvonget(clt *core.SDKClient, req *simba.TaobaosimbaadgrouponlineitemsvongetAPIRequest, session string) (*simba.TaobaosimbaadgrouponlineitemsvongetAPIResponse, error) {
	var resp simba.TaobaosimbaadgrouponlineitemsvongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
