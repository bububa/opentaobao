package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreDeliverdetailGetAPIResponse
查询发货单详情 API返回值
alibaba.nlife.store.deliverdetail.get

查询发货单详情 */
type AlibabaNlifeStoreDeliverdetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeStoreDeliverdetailGetAPIResponseModel
}

// AlibabaNlifeStoreDeliverdetailGetAPIResponseModel is 查询发货单详情 成功返回结果
type AlibabaNlifeStoreDeliverdetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_store_deliverdetail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果数据
	Data *DeliverDetailDo `json:"data,omitempty" xml:"data,omitempty"`
}
