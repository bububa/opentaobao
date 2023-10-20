package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupmobilediscountdelete 批量删除adgroup的移动溢价
// taobao.simba.adgroup.mobilediscount.delete
//
// 批量删除adgroup的移动溢价
func Taobaosimbaadgroupmobilediscountdelete(clt *core.SDKClient, req *simba.TaobaosimbaadgroupmobilediscountdeleteAPIRequest, session string) (*simba.TaobaosimbaadgroupmobilediscountdeleteAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupmobilediscountdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
