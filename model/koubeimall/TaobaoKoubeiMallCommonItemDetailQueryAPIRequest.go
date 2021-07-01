package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonItemDetailQueryAPIRequest
查询商品详情信息 API请求
taobao.koubei.mall.common.item.detail.query

查询口碑综合体内商品详情信息 */
type TaobaoKoubeiMallCommonItemDetailQueryAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈ID
	_mallId string
	// 门店ID
	_storeId string
	// 商品ID
	_itemId string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 终端设备描述(中、英文均可)
	_terminalType string
}

// New
