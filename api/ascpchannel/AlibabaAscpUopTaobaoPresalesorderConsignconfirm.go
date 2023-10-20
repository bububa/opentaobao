package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuoptaobaopresalesorderconsignconfirm 预售商家仓出库
// alibaba.ascp.uop.taobao.presalesorder.consignconfirm
//
// 预售商家仓出库
func Alibabaascpuoptaobaopresalesorderconsignconfirm(clt *core.SDKClient, req *ascpchannel.AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest, session string) (*ascpchannel.AlibabaascpuoptaobaopresalesorderconsignconfirmAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuoptaobaopresalesorderconsignconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
