package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziacluserrolesrevoke 回收账户被授予的角色接口
// alibaba.mozi.acl.userroles.revoke
//
// 调用此接口，会根据入参回收该账户下的该批角色
func Alibabamoziacluserrolesrevoke(clt *core.SDKClient, req *moziacl.AlibabamoziacluserrolesrevokeAPIRequest, session string) (*moziacl.AlibabamoziacluserrolesrevokeAPIResponse, error) {
	var resp moziacl.AlibabamoziacluserrolesrevokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
