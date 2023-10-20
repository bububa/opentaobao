package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoDataLogisticsDeliveryAgingPredict 配送物流时效预测
// cainiao.data.logistics.delivery.aging.predict
//
// 时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。
//
// 日常，展示具体的预测时效数值
//
// 大促期间，展示预测的时效区间
func CainiaoDataLogisticsDeliveryAgingPredict(clt *core.SDKClient, req *logistic.CainiaoDataLogisticsDeliveryAgingPredictAPIRequest, session string) (*logistic.CainiaoDataLogisticsDeliveryAgingPredictAPIResponse, error) {
	var resp logistic.CainiaoDataLogisticsDeliveryAgingPredictAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
