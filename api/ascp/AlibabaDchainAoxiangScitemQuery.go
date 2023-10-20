package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangscitemquery 货品查询
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
func Alibabadchainaoxiangscitemquery(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangscitemqueryAPIRequest, session string) (*ascp.AlibabadchainaoxiangscitemqueryAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangscitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
