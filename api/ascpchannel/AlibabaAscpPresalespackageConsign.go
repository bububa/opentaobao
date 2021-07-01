package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

/* AlibabaAscpPresalespackageConsign
预售预包尾款推单发货
alibaba.ascp.presalespackage.consign

预售预包尾款发货后推单处理 */
func AlibabaAscpPresalespackageConsign(clt *core.SDKClient, req *ascpchannel.AlibabaAscpPresalespackageConsignAPIRequest, session string) (*ascpchannel.AlibabaAscpPresalespackageConsignAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpPresalespackageConsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
