package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterservicestoredeleteservicestorecapacity 删除网点容量
// tmall.servicecenter.servicestore.deleteservicestorecapacity
//
// 删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
// 必选字段：serviceStoreCode、bizType
func Tmallservicecenterservicestoredeleteservicestorecapacity(clt *core.SDKClient, req *tmallsc.TmallservicecenterservicestoredeleteservicestorecapacityAPIRequest, session string) (*tmallsc.TmallservicecenterservicestoredeleteservicestorecapacityAPIResponse, error) {
	var resp tmallsc.TmallservicecenterservicestoredeleteservicestorecapacityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
