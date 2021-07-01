package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsSkuCreateAPIRequest
商品同步 API请求
taobao.wlb.wms.sku.create

商品同步 */
type TaobaoWlbWmsSkuCreateAPIRequest struct {
	model.Params
	// 商家商品编码
	_itemCode string
	// 条形码，多条码请用”；”分隔；
	_barCode string
	// 仓库编码
	_storeCode string
	// 商品名称
	_name string
	// 商品标题
	_title string
	// 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
	_type string
	// 商品类别编码（外部系统类别）
	_category string
	// 商品类别名称
	_categoryName string
	// 品牌编码
	_brand string
	// 品牌名称
	_brandName string
	// 规格
	_specification string
	// 颜色
	_color string
	// 尺码
	_size string
	// 毛重，单位克
	_grossWeight int64
	// 净重，单位克
	_netWeight int64
	// 长度，单位毫米
	_length int64
	// 宽度，单位毫米
	_width int64
	// 高度，单位毫米
	_height int64
	// 体积，单位立方厘米
	_volume int64
	// 箱规
	_pcs int64
	// 产地
	_originAddress int64
	// 批准文号
	_approvalNumber string
	// 是否启用保质期管理
	_isShelflife bool
	// 商品保质期天数
	_lifecycle int64
	// 保质期禁收天数
	_rejectLifecycle int64
	// 保质期禁售天数
	_lockupLifecycle int64
	// 保质期预警天数
	_adventLifecycle int64
	// 是否启用序列号管理
	_isSnMgt bool
	// 是否易碎品
	_isHygroscopic bool
	// 是否危险品
	_isDanger bool
	// 吊牌价，单位分
	_tagPrice int64
	// 零售价，单位分
	_itemPrice int64
	// 成本价，单位分
	_costPrice int64
	// 是否启用批次管理
	_isBatchMgt bool
	// 启用标识
	_useYn bool
	// 拓展属性
	_extendFields string
	// 商家商品ID
	_itemId string
	// 是否区域销售
	_isAreaSale bool
}

// New
