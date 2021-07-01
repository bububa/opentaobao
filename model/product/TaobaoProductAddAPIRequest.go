package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoProductAddAPIRequest
上传一个产品，不包括产品非主图和属性图片 API请求
taobao.product.add

获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 <br/>传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,<br/>调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.<br/>新增：套装产品发布，目前支持单件多个即 A*2 形式的套装 */
type TaobaoProductAddAPIRequest struct {
	model.Params
	// native_unkeyprops
	_nativeUnkeyprops string
	// 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id.
	_cid int64
	// 外部产品ID
	_outerId string
	// 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.
	_props string
	// 非关键属性结构:pid:vid;pid:vid.<br>非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font>
	_binds string
	// 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid.
	_saleProps string
	// 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”<br><font color=red>注：包含所有自定义属性的传入</font>
	_customerProps string
	// 产品市场价.精确到2位小数;单位为元.如：200.07
	_price string
	// 产品主图片.最大1M,目前仅支持GIF,JPG.
	_image *model.File
	// 产品名称,最大30个字符.
	_name string
	// 产品描述.最大不超过25000个字符
	_desc string
	// 是不是主图
	_major bool
	// 上市时间。目前只支持鞋城类目传入此参数
	_marketTime string
	// 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸
	_propertyAlias string
}

// New
