package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupmobilediscountupdate 对推广组进行单独移动溢价
// taobao.simba.adgroup.mobilediscount.update
//
// 对推广组进行单独移动溢价
func Taobaosimbaadgroupmobilediscountupdate(clt *core.SDKClient, req *simba.TaobaosimbaadgroupmobilediscountupdateAPIRequest, session string) (*simba.TaobaosimbaadgroupmobilediscountupdateAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupmobilediscountupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
