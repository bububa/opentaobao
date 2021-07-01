package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSellercatsListAddAPIRequest
添加卖家自定义类目 API请求
taobao.sellercats.list.add

此API添加卖家店铺内自定义类目 <br/>父类目parent_cid值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目 <br/>注：因为缓存的关系,添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布 */
type TaobaoSellercatsListAddAPIRequest struct {
	model.Params
	// 卖家自定义类目名称。不超过20个字符
	_name string
	// 链接图片URL地址。(绝对地址，格式：http://host/image_path)
	_pictUrl string
	// 父类目编号，如果类目为店铺下的一级类目：值等于0，如果类目为子类目，调用获取taobao.sellercats.list.get父类目编号
	_parentCid int64
	// 该类目在页面上的排序位置,取值范围:大于零的整数
	_sortOrder int64
}

// New
