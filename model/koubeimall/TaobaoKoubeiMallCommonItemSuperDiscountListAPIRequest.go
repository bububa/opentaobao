package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest
查询商圈内的超值特惠商品信息 API请求
taobao.koubei.mall.common.item.super.discount.list

查询商圈超值特惠商品信息列表 */
type TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈ID
	_mallId string
	// 查询商品最大个数，最大值50
	_itemSize int64
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
}

// New
