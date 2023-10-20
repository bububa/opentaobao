package cityretail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCityretailWmflConvertWarehouseAPIResponse 同城零售完美履约转仓 API返回值
// taobao.cityretail.wmfl.convert.warehouse
//
// 同城零售完美履约转仓
type TaobaoCityretailWmflConvertWarehouseAPIResponse struct {
	model.CommonResponse
	TaobaoCityretailWmflConvertWarehouseAPIResponseModel
}

// TaobaoCityretailWmflConvertWarehouseAPIResponseModel is 同城零售完美履约转仓 成功返回结果
type TaobaoCityretailWmflConvertWarehouseAPIResponseModel struct {
	XMLName xml.Name `xml:"cityretail_wmfl_convert_warehouse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *WorkResult `json:"result,omitempty" xml:"result,omitempty"`
}
