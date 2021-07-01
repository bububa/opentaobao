package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSellercatsListUpdateAPIRequest
更新卖家自定义类目 API请求
taobao.sellercats.list.update

此API更新卖家店铺内自定义类目 <br/>注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布 */
type TaobaoSellercatsListUpdateAPIRequest struct {
	model.Params
	// 卖家自定义类目编号
	_cid int64
	// 卖家自定义类目名称。不超过20个字符
	_name string
	// 链接图片URL地址
	_pictUrl string
	// 该类目在页面上的排序位置,取值范围:大于零的整数
	_sortOrder int64
}

// New
