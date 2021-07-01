package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

/* AlibabaAscpUopSupplierReverseorderInstorageResult
逆向销退入库单到仓结果回告
alibaba.ascp.uop.supplier.reverseorder.instorage.result

ERP回告销退入库单到仓信息回告 */
func AlibabaAscpUopSupplierReverseorderInstorageResult(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest, session string) (*ascpchannel.AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
