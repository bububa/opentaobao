package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiPriceGetAPIRequest
获取价格预估信息 API请求
alibaba.happytrip.taxi.price.get

打车价格预估 */
type AlibabaHappytripTaxiPriceGetAPIRequest struct {
	model.Params
	// 出发地纬度
	_flat string
	// 出发地经度
	_flng string
	// 目的地纬度
	_tlat string
	// 目的地经度
	_tlng string
	// 地图类型:amap：高德，默认高德地图
	_mapType string
	// 出发城市id
	_city string
	// 0:实时单 1:预约单
	_type int64
	// 预约单必须传（格式例如：2015-06-16 12:00:09）
	_departureTime string
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
	// 供应商车型代码
	_requireLevel string
	// 0：不拼车 1:允许拼车，默认不拼车
	_carpoolType int64
	// 乘车人数
	_passengerNumber int64
	// 用户唯一标识
	_uid string
	// 乘客手机号
	_passengerPhone string
}

// New
