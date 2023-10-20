package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupidsdeletedget 获取删除的推广组ID
// taobao.simba.adgroupids.deleted.get
//
// 获取删除的推广组ID
func Taobaosimbaadgroupidsdeletedget(clt *core.SDKClient, req *simba.TaobaosimbaadgroupidsdeletedgetAPIRequest, session string) (*simba.TaobaosimbaadgroupidsdeletedgetAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupidsdeletedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
