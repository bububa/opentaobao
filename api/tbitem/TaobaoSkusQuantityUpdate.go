package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoskusquantityupdate SKU库存修改
// taobao.skus.quantity.update
//
// 提供按照全量/增量的方式批量修改SKU库存的功能
func Taobaoskusquantityupdate(clt *core.SDKClient, req *tbitem.TaobaoskusquantityupdateAPIRequest, session string) (*tbitem.TaobaoskusquantityupdateAPIResponse, error) {
	var resp tbitem.TaobaoskusquantityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
