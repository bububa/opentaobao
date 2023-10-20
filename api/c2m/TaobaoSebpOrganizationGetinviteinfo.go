package c2m

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/c2m"
)

// Taobaosebporganizationgetinviteinfo 淘小铺机构上下级关系
// taobao.sebp.organization.getinviteinfo
//
// 机构人员获取机构上下级关系信息
func Taobaosebporganizationgetinviteinfo(clt *core.SDKClient, req *c2m.TaobaosebporganizationgetinviteinfoAPIRequest, session string) (*c2m.TaobaosebporganizationgetinviteinfoAPIResponse, error) {
	var resp c2m.TaobaosebporganizationgetinviteinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
