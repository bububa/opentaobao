package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoProductUpdateAPIRequest
修改一个产品，可以修改主图，不能修改子图片 API请求
taobao.product.update

传入产品ID <br/>可修改字段：outer_id,binds,sale_props,name,price,desc,image <br/>注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG<br/>      2.商城卖家产品发布24小时后不能作删除或修改操作 */
type TaobaoProductUpdateAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 外部产品ID
	_outerId string
	// 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
	_binds string
	// 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid
	_saleProps string
	// 产品市场价.精确到2位小数;单位为元.如:200.07
	_price string
	// 产品描述.最大不超过25000个字符
	_desc string
	// 产品主图.最大500K,目前仅支持GIF,JPG
	_image *model.File
	// 产品名称.最大不超过30个字符
	_name string
	// 是否是主图
	_major bool
	// 自定义非关键属性
	_nativeUnkeyprops string
}

// New
