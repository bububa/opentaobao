package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoDataLogisticsDeliveryAgingPredictAPIRequest
配送物流时效预测 API请求
cainiao.data.logistics.delivery.aging.predict

时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。

日常，展示具体的预测时效数值

大促期间，展示预测的时效区间 */
type CainiaoDataLogisticsDeliveryAgingPredictAPIRequest struct {
	model.Params
	// 发货城市
	_sendCityName string
	// 发货区
	_sendCountyName string
	// 发货详细地址
	_sendAddr string
	// 发货省
	_sendProvName string
	// 收货市
	_recCityName string
	// 收货详细地址
	_recAddr string
	// 收货区
	_recCountyName string
	// 收货省
	_recProvName string
	// 收货街道
	_recTownName string
}

// New
