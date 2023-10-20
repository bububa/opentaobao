package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupschangedget 分页获取修改的推广组ID和修改时间
// taobao.simba.adgroups.changed.get
//
// 分页获取修改的推广组ID和修改时间
func Taobaosimbaadgroupschangedget(clt *core.SDKClient, req *simba.TaobaosimbaadgroupschangedgetAPIRequest, session string) (*simba.TaobaosimbaadgroupschangedgetAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupschangedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
