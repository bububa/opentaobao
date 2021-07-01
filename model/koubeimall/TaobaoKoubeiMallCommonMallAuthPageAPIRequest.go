package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonMallAuthPageAPIRequest
分页查询已授权的商圈列表信息 API请求
taobao.koubei.mall.common.mall.auth.page

分页查询口碑已授权商圈的列表信息 */
type TaobaoKoubeiMallCommonMallAuthPageAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 分页查询起始值，默认为0
	_start int64
	// 每页查询量，默认10（建议查询值为10的倍数，最大不超过30）
	_pageSize int64
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
}

// New
