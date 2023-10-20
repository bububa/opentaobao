package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpglobalsupplieritemlistinfoquery 供应商供品信息查询
// alibaba.ascp.global.supplier.item.list.info.query
//
// 供应商供品信息查询
func Alibabaascpglobalsupplieritemlistinfoquery(clt *core.SDKClient, req *ascpchannel.AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest, session string) (*ascpchannel.AlibabaascpglobalsupplieritemlistinfoqueryAPIResponse, error) {
	var resp ascpchannel.AlibabaascpglobalsupplieritemlistinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
