package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangscitemqueryAPIResponse 货品查询 API返回值
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
type AlibabadchainaoxiangscitemqueryAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangscitemqueryAPIResponseModel
}

// AlibabadchainaoxiangscitemqueryAPIResponseModel is 货品查询 成功返回结果
type AlibabadchainaoxiangscitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	QueryScitemResponse *TopResponse `json:"query_scitem_response,omitempty" xml:"query_scitem_response,omitempty"`
}
