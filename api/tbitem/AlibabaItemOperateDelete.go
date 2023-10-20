package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitemoperatedelete 商品删除
// alibaba.item.operate.delete
//
// 商品删除
func Alibabaitemoperatedelete(clt *core.SDKClient, req *tbitem.AlibabaitemoperatedeleteAPIRequest, session string) (*tbitem.AlibabaitemoperatedeleteAPIResponse, error) {
	var resp tbitem.AlibabaitemoperatedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
