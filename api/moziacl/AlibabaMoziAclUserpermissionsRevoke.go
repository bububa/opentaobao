package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziacluserpermissionsrevoke 回收账户权限
// alibaba.mozi.acl.userpermissions.revoke
//
// 调用此接口，会根据入参回收该账户下的该批权限点
func Alibabamoziacluserpermissionsrevoke(clt *core.SDKClient, req *moziacl.AlibabamoziacluserpermissionsrevokeAPIRequest, session string) (*moziacl.AlibabamoziacluserpermissionsrevokeAPIResponse, error) {
	var resp moziacl.AlibabamoziacluserpermissionsrevokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
