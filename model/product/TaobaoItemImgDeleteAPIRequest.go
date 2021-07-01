package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemImgDeleteAPIRequest
删除商品图片 API请求
taobao.item.img.delete

删除商品图片 */
type TaobaoItemImgDeleteAPIRequest struct {
	model.Params
	// 商品数字ID
	_numIid int64
	// 商品图片ID；如果是竖图，请将id的值设置为1
	_id int64
	// 标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下
	_isSixthPic bool
}

// New
