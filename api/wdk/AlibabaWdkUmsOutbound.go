package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkUmsOutbound
出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
alibaba.wdk.ums.outbound

出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) */
func AlibabaWdkUmsOutbound(clt *core.SDKClient, req *wdk.AlibabaWdkUmsOutboundAPIRequest, session string) (*wdk.AlibabaWdkUmsOutboundAPIResponse, error) {
	var resp wdk.AlibabaWdkUmsOutboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
