package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemUpdateDelistingAPIRequest
商品下架 API请求
taobao.item.update.delisting

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateDelistingAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// New
