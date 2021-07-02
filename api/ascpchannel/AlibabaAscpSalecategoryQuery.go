package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpSalecategoryQuery 货品品类查询
// alibaba.ascp.salecategory.query
//
// 根据货品ID查询对应销售品类ID
func AlibabaAscpSalecategoryQuery(clt *core.SDKClient, req *ascpchannel.AlibabaAscpSalecategoryQueryAPIRequest, session string) (*ascpchannel.AlibabaAscpSalecategoryQueryAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpSalecategoryQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
