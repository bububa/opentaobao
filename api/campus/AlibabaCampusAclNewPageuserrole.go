package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewpageuserrole 分页查询管理员
// alibaba.campus.acl.new.pageuserrole
//
// 新增用户和角色的关系
func Alibabacampusaclnewpageuserrole(clt *core.SDKClient, req *campus.AlibabacampusaclnewpageuserroleAPIRequest, session string) (*campus.AlibabacampusaclnewpageuserroleAPIResponse, error) {
	var resp campus.AlibabacampusaclnewpageuserroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
