package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonStorePageAPIRequest
分页查询综合体内的门店列表信息 API请求
taobao.koubei.mall.common.store.page

分页查询综合体内的门店列表信息 */
type TaobaoKoubeiMallCommonStorePageAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈ID
	_mallId string
	// 分页查询起始值，默认为0
	_start int64
	// 每页查询量，默认10（建议查询值为10的倍数，最大不超过20）
	_pageSize int64
	// 门店列表按照类目筛选条件过滤，可通过查询商圈详情获取类目信息
	_categoryIds []string
	// 商圈内的门店ID
	_storeId string
	// 门店列表排序规则；默认：门店质量分降序，暂无其它排序规则
	_order string
	// 店铺服务标签，用于列表过滤条件；比如：点餐/外卖/预定等服务筛选条件。预定：SERVICE_DING；排号：SERVICE_PAI；点菜：SERVICE_DIAN；外卖：SERVICE_WAI；
	_serviceTag []string
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
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
}

// New
