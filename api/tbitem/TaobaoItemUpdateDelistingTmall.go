package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemUpdateDelistingTmall taobao.item.update.delisting.tmall
// taobao.item.update.delisting.tmall
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
func TaobaoItemUpdateDelistingTmall(clt *core.SDKClient, req *tbitem.TaobaoItemUpdateDelistingTmallAPIRequest, session string) (*tbitem.TaobaoItemUpdateDelistingTmallAPIResponse, error) {
	var resp tbitem.TaobaoItemUpdateDelistingTmallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
