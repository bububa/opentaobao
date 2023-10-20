package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopselfsupplierwaybillquery 商家仓自营配电子面单取号
// alibaba.ascp.uop.self.supplier.waybill.query
//
// ERP调用打印面单取号接口
func Alibabaascpuopselfsupplierwaybillquery(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopselfsupplierwaybillqueryAPIRequest, session string) (*ascpchannel.AlibabaascpuopselfsupplierwaybillqueryAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopselfsupplierwaybillqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
