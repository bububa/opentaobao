package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitemoperatedownshelf 商品下架
// alibaba.item.operate.downshelf
//
// 商品下架
func Alibabaitemoperatedownshelf(clt *core.SDKClient, req *tbitem.AlibabaitemoperatedownshelfAPIRequest, session string) (*tbitem.AlibabaitemoperatedownshelfAPIResponse, error) {
	var resp tbitem.AlibabaitemoperatedownshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
