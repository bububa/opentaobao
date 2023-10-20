package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemimgdelete 删除商品图片
// taobao.item.img.delete
//
// 删除商品图片
func Taobaoitemimgdelete(clt *core.SDKClient, req *tbitem.TaobaoitemimgdeleteAPIRequest, session string) (*tbitem.TaobaoitemimgdeleteAPIResponse, error) {
	var resp tbitem.TaobaoitemimgdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
