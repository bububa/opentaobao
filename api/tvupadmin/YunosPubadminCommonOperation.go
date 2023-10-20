package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosPubadminCommonOperation 内部迎客松通用服务
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
func YunosPubadminCommonOperation(clt *core.SDKClient, req *tvupadmin.YunosPubadminCommonOperationAPIRequest, resp *tvupadmin.YunosPubadminCommonOperationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
