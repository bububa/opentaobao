package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemUpdateAPIRequest
根据商品ID或商家编码修改后端商品 API请求
taobao.scitem.update

根据商品ID或商家编码修改后端商品 */
type TaobaoScitemUpdateAPIRequest struct {
	model.Params
	// 后端商品ID，跟outer_code必须指定一个
	_itemId int64
	// 商家编码，跟item_id必须指定一个
	_outerCode string
	// 商品名称
	_itemName string
	// 0.普通供应链商品 1.供应链组合主商品
	_itemType int64
	// 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
	_updateProperties string
	// 条形码
	_barCode string
	// 仓储商编码
	_wmsCode string
	// 是否易碎 0：不是  1：是
	_isFriable int64
	// 是否危险 0：不是  0：是
	_isDangerous int64
	// 是否是贵重品 0:不是 1：是
	_isCostly int64
	// 是否保质期：0:不是 1：是
	_isWarranty int64
	// weight
	_weight int64
	// 长度 单位：mm
	_length int64
	// 宽 单位：mm
	_width int64
	// 高 单位：mm
	_height int64
	// 体积：立方厘米
	_volume int64
	// price
	_price int64
	// remark
	_remark string
	// 0:液体，1：粉体，2：固体
	_matterStatus int64
	// 品牌id
	_brandId int64
	// brand_Name
	_brandName string
	// 淘宝SKU产品级编码CSPU ID
	_spuId int64
	// 移除商品属性P列表,P由系统分配：p1；p2
	_removeProperties string
	// 1表示区域销售，0或是空是非区域销售
	_isAreaSale int64
}

// New
