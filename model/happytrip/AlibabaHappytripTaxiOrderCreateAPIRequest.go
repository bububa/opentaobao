package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderCreateAPIRequest
用户叫车 API请求
alibaba.happytrip.taxi.order.create

用户根据需要发起叫车请求，在发起请求之前必须事先获得order id. */
type AlibabaHappytripTaxiOrderCreateAPIRequest struct {
	model.Params
	// 用户唯一标识
	_uid string
	// 请求id 获取请参见
	_orderId string
	// 叫车车型，0(实时)；1(预约)
	_type int64
	// 乘客手机号
	_passengerPhone string
	// 出发地城市
	_city string
	// 出发地纬度
	_flat string
	// 出发地经度
	_flng string
	// 出发地名称(最多50个字)
	_startName string
	// 出发地详细地址(最多100个字)
	_startAddress string
	// 目的地纬度
	_tlat string
	// 目的地经度
	_tlng string
	// 目的地名称(最多50个字)
	_endName string
	// 目的地详细地址(最多100个字)
	_endAddress string
	// 当前位置纬度
	_clat string
	// 当前位置经度
	_clng string
	// 出发时间，不传表示现在用车（例如：2015-06-16 12:00:09）
	_departureTime string
	// 车型代码
	_requireLevel string
	// 客户端时间（例如：2015-06-16 12:00:09）
	_appTime string
	// 地图类型:amap：高德，默认高德地图
	_mapType string
	// 发送短信策略(0,为叫车人和乘车人都发送，1，乘车人发送叫车人不发，2乘车人不发叫车人发，3乘车人和叫车人都不发)
	_smsPolicy int64
	// 备注
	_extraInfo string
	// 价格md5,通过 预估价接口获得
	_dynamicMd5 string
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
	// 线路类型，0或空表示普通线路；1，表示一口价路线
	_lineType int64
	// 0：不拼车 1:允许拼车，默认不拼车
	_carpoolType int64
	// 乘车人数，允许拼车时该参数不为空
	_passengerNumber int64
}

// New
