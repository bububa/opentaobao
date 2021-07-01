package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiMallCommonStoreCommentPageAPIRequest
分页查询门店评论详情信息 API请求
taobao.koubei.mall.common.store.comment.page

查询口碑综合体内的门店评论信息 */
type TaobaoKoubeiMallCommonStoreCommentPageAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈ID
	_mallId string
	// 门店ID
	_storeId string
	// 查询起始值，默认为0
	_start int64
	// 每页查询量，默认为20，最大数值20
	_pageSize int64
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 纬度（终端设备地理位置）
	_longitude string
	// 经度（终端设备地理位置）
	_latitude string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
}

// New
