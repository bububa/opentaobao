package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsoutbound 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
// alibaba.wdk.ums.outbound
//
// 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
func Alibabawdkumsoutbound(clt *core.SDKClient, req *wdk.AlibabawdkumsoutboundAPIRequest, session string) (*wdk.AlibabawdkumsoutboundAPIResponse, error) {
	var resp wdk.AlibabawdkumsoutboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
