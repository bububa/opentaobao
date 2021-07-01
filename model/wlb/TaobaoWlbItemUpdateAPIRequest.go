package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemUpdateAPIRequest
物流宝商品修改 API请求
taobao.wlb.item.update

修改物流宝商品信息 */
type TaobaoWlbItemUpdateAPIRequest struct {
	model.Params
	// 需要修改的商品属性值的列表，如果属性不存在，则新增属性
	_updatePropertyKeyList string
	// 需要删除的商品属性key列表
	_deletePropertyKeyList string
	// 需要修改的属性值的列表
	_updatePropertyValueList string
	// 要修改的商品id
	_id int64
	// 要修改的商品名称
	_name string
	// 要修改的商品标题
	_title string
	// 要修改的商品备注
	_remark string
	// 是否易碎品
	_isFriable bool
	// 是否危险品
	_isDangerous bool
	// 商品颜色
	_color string
	// 商品重量，单位G
	_weight int64
	// 商品长度，单位厘米
	_length int64
	// 商品宽度，单位厘米
	_width int64
	// 商品高度，单位厘米
	_height int64
	// 商品体积，单位立方厘米
	_volume int64
	// 商品货类
	_goodsCat string
	// 商品计价货类
	_pricingCat string
	// 商品包装材料类型
	_packageMaterial string
}

// New
