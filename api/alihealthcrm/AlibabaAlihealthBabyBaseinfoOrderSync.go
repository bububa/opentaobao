package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthbabybaseinfoordersync alibaba.alihealth.baby.baseinfo.order.sync
// alibaba.alihealth.baby.baseinfo.order.sync
//
// 育学园将订单信息回传给我们
func Alibabaalihealthbabybaseinfoordersync(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthbabybaseinfoordersyncAPIRequest, session string) (*alihealthcrm.AlibabaalihealthbabybaseinfoordersyncAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthbabybaseinfoordersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
