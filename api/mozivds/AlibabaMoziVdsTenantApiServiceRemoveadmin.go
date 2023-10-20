package mozivds

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozivds"
)

// Alibabamozivdstenantapiserviceremoveadmin 删除租户管理员服务
// alibaba.mozi.vds.tenant.api.service.removeadmin
//
// 删除租户管理员top服务
func Alibabamozivdstenantapiserviceremoveadmin(clt *core.SDKClient, req *mozivds.AlibabamozivdstenantapiserviceremoveadminAPIRequest, session string) (*mozivds.AlibabamozivdstenantapiserviceremoveadminAPIResponse, error) {
	var resp mozivds.AlibabamozivdstenantapiserviceremoveadminAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
