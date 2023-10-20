package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupdelete 删除一个推广组
// taobao.simba.adgroup.delete
//
// 删除一个推广组
func Taobaosimbaadgroupdelete(clt *core.SDKClient, req *simba.TaobaosimbaadgroupdeleteAPIRequest, session string) (*simba.TaobaosimbaadgroupdeleteAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
