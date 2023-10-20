package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicedeviceorderupdate 回传/更新设备订购单
// alibaba.einvoice.device.order.update
//
// 更新设备订购单，同步税控设备信息
func Alibabaeinvoicedeviceorderupdate(clt *core.SDKClient, req *einvoice.AlibabaeinvoicedeviceorderupdateAPIRequest, session string) (*einvoice.AlibabaeinvoicedeviceorderupdateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicedeviceorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
