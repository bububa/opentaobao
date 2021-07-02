package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoDataLogisticsCpDeliveryAgingPredict CP配送物流时效预测
// cainiao.data.logistics.cp.delivery.aging.predict
//
// 时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。
//
// 日常时效是数值字符串
// 大促时效是数值区间字符串
// 方式1:
// 输入：发货省、市、区、详细地址，收货省、市、区、街道、详细地址，快递公司ID
// 输出：预估时效（小时数）
func CainiaoDataLogisticsCpDeliveryAgingPredict(clt *core.SDKClient, req *logistic.CainiaoDataLogisticsCpDeliveryAgingPredictAPIRequest, session string) (*logistic.CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse, error) {
	var resp logistic.CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
