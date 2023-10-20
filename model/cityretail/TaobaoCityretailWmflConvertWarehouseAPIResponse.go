package cityretail

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoCityretailWmflConvertWarehouseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCityretailWmflConvertWarehouseAPIResponseModel).Reset()
}

// TaobaoCityretailWmflConvertWarehouseAPIResponseModel is 同城零售完美履约转仓 成功返回结果
type TaobaoCityretailWmflConvertWarehouseAPIResponseModel struct {
	XMLName xml.Name `xml:"cityretail_wmfl_convert_warehouse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *WorkResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCityretailWmflConvertWarehouseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoCityretailWmflConvertWarehouseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCityretailWmflConvertWarehouseAPIResponse)
	},
}

// GetTaobaoCityretailWmflConvertWarehouseAPIResponse 从 sync.Pool 获取 TaobaoCityretailWmflConvertWarehouseAPIResponse
func GetTaobaoCityretailWmflConvertWarehouseAPIResponse() *TaobaoCityretailWmflConvertWarehouseAPIResponse {
	return poolTaobaoCityretailWmflConvertWarehouseAPIResponse.Get().(*TaobaoCityretailWmflConvertWarehouseAPIResponse)
}

// ReleaseTaobaoCityretailWmflConvertWarehouseAPIResponse 将 TaobaoCityretailWmflConvertWarehouseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCityretailWmflConvertWarehouseAPIResponse(v *TaobaoCityretailWmflConvertWarehouseAPIResponse) {
	v.Reset()
	poolTaobaoCityretailWmflConvertWarehouseAPIResponse.Put(v)
}
