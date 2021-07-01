package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonMallNearListAPIRequest
根据POI查询附近商圈列表信息 API请求
taobao.koubei.mall.common.mall.near.list

通过用户/终端设备地理位置POI信息，查询附近商圈信息 */
type TaobaoKoubeiMallCommonMallNearListAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 召回半径，单位m，最大数值不能超过10km（该字段为空，默认全城召回）
	_radius int64
	// 查询个数，最大查询量不能超过50个
	_size int64
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 终端设备描述(中、英文均可)
	_terminalType string
}

// New
