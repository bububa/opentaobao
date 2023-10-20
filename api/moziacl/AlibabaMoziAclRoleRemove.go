package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziaclroleremove 删除角色
// alibaba.mozi.acl.role.remove
//
// 根据传入的角色code、租户id，删除租户内对应的角色
func Alibabamoziaclroleremove(clt *core.SDKClient, req *moziacl.AlibabamoziaclroleremoveAPIRequest, session string) (*moziacl.AlibabamoziaclroleremoveAPIResponse, error) {
	var resp moziacl.AlibabamoziaclroleremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
