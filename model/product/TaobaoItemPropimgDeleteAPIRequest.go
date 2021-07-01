package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemPropimgDeleteAPIRequest
删除属性图片 API请求
taobao.item.propimg.delete

删除propimg_id 所指定的商品属性图片 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>propimg_id对应的属性图片需要属于num_iid对应的商品 */
type TaobaoItemPropimgDeleteAPIRequest struct {
	model.Params
	// 商品属性图片ID
	_id int64
	// 商品数字ID，必选
	_numIid int64
}

// New
