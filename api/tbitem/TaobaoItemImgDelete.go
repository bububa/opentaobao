package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemImgDelete 删除商品图片
// taobao.item.img.delete
//
// 删除商品图片
func TaobaoItemImgDelete(clt *core.SDKClient, req *tbitem.TaobaoItemImgDeleteAPIRequest, session string) (*tbitem.TaobaoItemImgDeleteAPIResponse, error) {
	var resp tbitem.TaobaoItemImgDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
