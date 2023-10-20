package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// Alibabaretaildeviceroadstatusreset 贩卖机货道解锁
// alibaba.retail.device.road.status.reset
//
// 贩卖机货道解锁
func Alibabaretaildeviceroadstatusreset(clt *core.SDKClient, req *retail.AlibabaretaildeviceroadstatusresetAPIRequest, session string) (*retail.AlibabaretaildeviceroadstatusresetAPIResponse, error) {
	var resp retail.AlibabaretaildeviceroadstatusresetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
