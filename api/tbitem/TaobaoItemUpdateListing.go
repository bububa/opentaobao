package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemupdatelisting 一口价商品上架
// taobao.item.update.listing
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
func Taobaoitemupdatelisting(clt *core.SDKClient, req *tbitem.TaobaoitemupdatelistingAPIRequest, session string) (*tbitem.TaobaoitemupdatelistingAPIResponse, error) {
	var resp tbitem.TaobaoitemupdatelistingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
