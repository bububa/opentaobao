package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest
CP配送物流时效预测 API请求
cainiao.data.logistics.cp.delivery.aging.predict

时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。

日常时效是数值字符串
大促时效是数值区间字符串
方式1:
输入：发货省、市、区、详细地址，收货省、市、区、街道、详细地址，快递公司ID
输出：预估时效（小时数） */
type CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest struct {
	model.Params
	// 发货城市
	_sendCityName string
	// 发货区
	_sendCountyName string
	// 自己输入的详细发货地址
	_sendAddr string
	// 发货省
	_sendProvName string
	// 收货城市
	_recCityName string
	// 自己输入的详细收货地址
	_recAddr string
	// 收货区
	_recCountyName string
	// 收货省
	_recProvName string
	// 第四级，一般是收货街道
	_recTownName string
	// 物流公司id
	_cpId string
}

// New
