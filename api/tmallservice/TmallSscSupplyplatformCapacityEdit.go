package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallSscSupplyplatformCapacityEdit 容量编辑
// tmall.ssc.supplyplatform.capacity.edit
//
// 容量编辑
func TmallSscSupplyplatformCapacityEdit(clt *core.SDKClient, req *tmallservice.TmallSscSupplyplatformCapacityEditAPIRequest, resp *tmallservice.TmallSscSupplyplatformCapacityEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
