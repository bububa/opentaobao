package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaodatalogisticsdeliveryagingpredictAPIResponse 配送物流时效预测 API返回值
// cainiao.data.logistics.delivery.aging.predict
//
// 时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。
//
// 日常，展示具体的预测时效数值
//
// 大促期间，展示预测的时效区间
type CainiaodatalogisticsdeliveryagingpredictAPIResponse struct {
	model.CommonResponse
	CainiaodatalogisticsdeliveryagingpredictAPIResponseModel
}

// CainiaodatalogisticsdeliveryagingpredictAPIResponseModel is 配送物流时效预测 成功返回结果
type CainiaodatalogisticsdeliveryagingpredictAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_data_logistics_delivery_aging_predict_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 列表类型返回
	Datas []CainiaodatalogisticsdeliveryagingpredictData `json:"datas,omitempty" xml:"datas>cainiaodatalogisticsdeliveryagingpredict_data,omitempty"`
}
