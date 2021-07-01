package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosPubadminCommonOperation
内部迎客松通用服务
yunos.pubadmin.common.operation

内部迎客松通用服务 */
func YunosPubadminCommonOperation(clt *core.SDKClient, req *tvupadmin.YunosPubadminCommonOperationAPIRequest, session string) (*tvupadmin.YunosPubadminCommonOperationAPIResponse, error) {
	var resp tvupadmin.YunosPubadminCommonOperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
