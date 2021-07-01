package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

/* AlibabaAscpUopTobPackageQuery
B2B包裹查询接口
alibaba.ascp.uop.tob.package.query

供应链中台TOB包裹查询接口 */
func AlibabaAscpUopTobPackageQuery(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopTobPackageQueryAPIRequest, session string) (*ascpchannel.AlibabaAscpUopTobPackageQueryAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpUopTobPackageQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
