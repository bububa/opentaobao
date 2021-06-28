package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
配送物流时效预测 APIResponse
cainiao.data.logistics.delivery.aging.predict

时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。

日常，展示具体的预测时效数值

大促期间，展示预测的时效区间
*/
type CainiaoDataLogisticsDeliveryAgingPredictAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_data_logistics_delivery_aging_predict_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 列表类型返回
    
    Datas   []CainiaoDataLogisticsDeliveryAgingPredictData `json:"datas,omitempty" xml:"