package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemUpdateDelistingTmallAPIRequest
taobao.item.update.delisting.tmall API请求
taobao.item.update.delisting.tmall

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateDelistingTmallAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// New
