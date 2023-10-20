package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemupdatedelistingtmall taobao.item.update.delisting.tmall
// taobao.item.update.delisting.tmall
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
func Taobaoitemupdatedelistingtmall(clt *core.SDKClient, req *tbitem.TaobaoitemupdatedelistingtmallAPIRequest, session string) (*tbitem.TaobaoitemupdatedelistingtmallAPIResponse, error) {
	var resp tbitem.TaobaoitemupdatedelistingtmallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
