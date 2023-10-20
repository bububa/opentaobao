package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemupdatedelisting 商品下架
// taobao.item.update.delisting
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
func Taobaoitemupdatedelisting(clt *core.SDKClient, req *tbitem.TaobaoitemupdatedelistingAPIRequest, session string) (*tbitem.TaobaoitemupdatedelistingAPIResponse, error) {
	var resp tbitem.TaobaoitemupdatedelistingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
