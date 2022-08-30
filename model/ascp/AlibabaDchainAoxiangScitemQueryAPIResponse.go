package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemQueryAPIResponse 货品查询 API返回值
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
type AlibabaDchainAoxiangScitemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangScitemQueryAPIResponseModel
}

// AlibabaDchainAoxiangScitemQueryAPIResponseModel is 货品查询 成功返回结果
type AlibabaDchainAoxiangScitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	QueryScitemResponse *QueryScItemResponse `json:"query_scitem_response,omitempty" xml:"query_scitem_response,omitempty"`
}
