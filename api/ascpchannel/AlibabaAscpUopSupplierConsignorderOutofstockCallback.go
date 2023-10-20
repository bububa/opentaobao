package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopsupplierconsignorderoutofstockcallback 履约单纬度的仓缺货回告服务
// alibaba.ascp.uop.supplier.consignorder.outofstock.callback
//
// 商家仓履约单纬度的仓缺货回告接口
func Alibabaascpuopsupplierconsignorderoutofstockcallback(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest, session string) (*ascpchannel.AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
