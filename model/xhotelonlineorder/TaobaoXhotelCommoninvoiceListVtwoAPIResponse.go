package xhotelonlineorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcommoninvoicelistvtwoAPIResponse 用户常用发票信息查询接口 API返回值
// taobao.xhotel.commoninvoice.list.vtwo
//
// 获取用户常用发票信息接口
type TaobaoxhotelcommoninvoicelistvtwoAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelcommoninvoicelistvtwoAPIResponseModel
}

// TaobaoxhotelcommoninvoicelistvtwoAPIResponseModel is 用户常用发票信息查询接口 成功返回结果
type TaobaoxhotelcommoninvoicelistvtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_commoninvoice_list_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoxhotelcommoninvoicelistvtwoResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
