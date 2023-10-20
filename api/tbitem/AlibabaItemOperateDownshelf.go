package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemOperateDownshelf 商品下架
// alibaba.item.operate.downshelf
//
// 商品下架
func AlibabaItemOperateDownshelf(clt *core.SDKClient, req *tbitem.AlibabaItemOperateDownshelfAPIRequest, session string) (*tbitem.AlibabaItemOperateDownshelfAPIResponse, error) {
	var resp tbitem.AlibabaItemOperateDownshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
