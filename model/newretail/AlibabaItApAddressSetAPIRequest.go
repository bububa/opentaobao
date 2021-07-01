package newretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItApAddressSetAPIRequest
setApAddressNew API请求
alibaba.it.ap.address.set

该接口可为ISV系统提供 ap位置信息维护的功能 */
type AlibabaItApAddressSetAPIRequest struct {
	model.Params
	// 城市
	_apCityName string
	// 经度
	_lng string
	// 签名
	_signature string
	// 园区/门店
	_apCampusName string
	// 区域
	_apAreaName string
	// 省份
	_apProvinceName string
	// ap mac地址
	_mac string
	// ap空间单元名称
	_apUnitName string
	// 楼层
	_apFloor string
	// 楼栋
	_apBuildingName string
	// 分配的内部ak
	_appKeyInternal string
	// 国家
	_apNationName string
	// 纬度
	_lat string
	// 方位
	_direction string
	// 时间戳，毫秒
	_timestampInternal int64
}

// New
