package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpsalecategoryqueryAPIResponse 货品品类查询 API返回值
// alibaba.ascp.salecategory.query
//
// 根据货品ID查询对应销售品类ID
type AlibabaascpsalecategoryqueryAPIResponse struct {
	model.CommonResponse
	AlibabaascpsalecategoryqueryAPIResponseModel
}

// AlibabaascpsalecategoryqueryAPIResponseModel is 货品品类查询 成功返回结果
type AlibabaascpsalecategoryqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_salecategory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 货品查询结构化对象
	DataList []SalecategoryQueryResponse `json:"data_list,omitempty" xml:"data_list>salecategory_query_response,omitempty"`
}
