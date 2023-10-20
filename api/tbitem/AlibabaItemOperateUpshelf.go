package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitemoperateupshelf 商品上架
// alibaba.item.operate.upshelf
//
// 商品上架
func Alibabaitemoperateupshelf(clt *core.SDKClient, req *tbitem.AlibabaitemoperateupshelfAPIRequest, session string) (*tbitem.AlibabaitemoperateupshelfAPIResponse, error) {
	var resp tbitem.AlibabaitemoperateupshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
