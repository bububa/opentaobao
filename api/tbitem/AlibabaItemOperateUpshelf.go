package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemOperateUpshelf 商品上架
// alibaba.item.operate.upshelf
//
// 商品上架
func AlibabaItemOperateUpshelf(clt *core.SDKClient, req *tbitem.AlibabaItemOperateUpshelfAPIRequest, session string) (*tbitem.AlibabaItemOperateUpshelfAPIResponse, error) {
	var resp tbitem.AlibabaItemOperateUpshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
