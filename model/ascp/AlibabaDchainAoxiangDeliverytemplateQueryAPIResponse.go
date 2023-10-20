package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangdeliverytemplatequeryAPIResponse 商家运费模板查询 API返回值
// alibaba.dchain.aoxiang.deliverytemplate.query
//
// 商家运费模板查询
type AlibabadchainaoxiangdeliverytemplatequeryAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangdeliverytemplatequeryAPIResponseModel
}

// AlibabadchainaoxiangdeliverytemplatequeryAPIResponseModel is 商家运费模板查询 成功返回结果
type AlibabadchainaoxiangdeliverytemplatequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_deliverytemplate_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	QueryDeliverytemplateResponse *TopResponse `json:"query_deliverytemplate_response,omitempty" xml:"query_deliverytemplate_response,omitempty"`
}
