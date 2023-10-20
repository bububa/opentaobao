package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliiquerybywaybillcodeAPIResponse 通过面单号查询面单打印报文 API返回值
// cainiao.waybill.ii.query.by.waybillcode
//
// 通过面单号查询面单的打印报文
type CainiaowaybilliiquerybywaybillcodeAPIResponse struct {
	model.CommonResponse
	CainiaowaybilliiquerybywaybillcodeAPIResponseModel
}

// CainiaowaybilliiquerybywaybillcodeAPIResponseModel is 通过面单号查询面单打印报文 成功返回结果
type CainiaowaybilliiquerybywaybillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_query_by_waybillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回值
	Modules []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_with_result_desc_response,omitempty"`
}
