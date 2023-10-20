package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupupdate 更新一个推广组的信息
// taobao.simba.adgroup.update
//
// 更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
func Taobaosimbaadgroupupdate(clt *core.SDKClient, req *simba.TaobaosimbaadgroupupdateAPIRequest, session string) (*simba.TaobaosimbaadgroupupdateAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
