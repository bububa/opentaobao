package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpsalecategoryquery 货品品类查询
// alibaba.ascp.salecategory.query
//
// 根据货品ID查询对应销售品类ID
func Alibabaascpsalecategoryquery(clt *core.SDKClient, req *ascpchannel.AlibabaascpsalecategoryqueryAPIRequest, session string) (*ascpchannel.AlibabaascpsalecategoryqueryAPIResponse, error) {
	var resp ascpchannel.AlibabaascpsalecategoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
