package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest
查询门店推荐菜信息 API请求
taobao.koubei.mall.common.store.display.goods.list

提供查询口碑商圈内的门店推荐菜信息 */
type TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 门店ID
	_storeId string
	// 商圈ID
	_mallId string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
}

// New
