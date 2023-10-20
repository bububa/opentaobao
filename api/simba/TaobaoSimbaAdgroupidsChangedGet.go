package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupidschangedget 获取修改的推广组ID
// taobao.simba.adgroupids.changed.get
//
// 获取修改的推广组ID
func Taobaosimbaadgroupidschangedget(clt *core.SDKClient, req *simba.TaobaosimbaadgroupidschangedgetAPIRequest, session string) (*simba.TaobaosimbaadgroupidschangedgetAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupidschangedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
