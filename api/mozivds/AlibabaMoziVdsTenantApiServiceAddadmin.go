package mozivds

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozivds"
)

// Alibabamozivdstenantapiserviceaddadmin 新建租户管理员
// alibaba.mozi.vds.tenant.api.service.addadmin
//
// 新建租户管理员
// alibaba.mozi.vds.tenant.api.service.addadmin
func Alibabamozivdstenantapiserviceaddadmin(clt *core.SDKClient, req *mozivds.AlibabamozivdstenantapiserviceaddadminAPIRequest, session string) (*mozivds.AlibabamozivdstenantapiserviceaddadminAPIResponse, error) {
	var resp mozivds.AlibabamozivdstenantapiserviceaddadminAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
