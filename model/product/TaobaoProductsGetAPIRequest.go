package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoProductsGetAPIRequest
获取产品列表 API请求
taobao.products.get

根据淘宝会员帐号搜索所有产品信息，推荐使用taobao.products.search
注意：支持分页，每页最多返回100条,默认值为40,页码从1开始，默认为第一页 */
type TaobaoProductsGetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔
	_fields []string
	// 用户昵称
	_nick string
	// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.
	_pageNo int64
	// 每页条数.每页返回最多返回100条,默认值为40
	_pageSize int64
}

// New
