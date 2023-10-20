package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse 虚拟供应商履约接口切换 API返回值
// taobao.virtual.dsf.supplier.interface.switch
//
// 虚拟供应商履约接口切换
type TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse struct {
	model.CommonResponse
	TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponseModel).Reset()
}

// TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponseModel is 虚拟供应商履约接口切换 成功返回结果
type TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponseModel struct {
	XMLName xml.Name `xml:"virtual_dsf_supplier_interface_switch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *DsfSupplierSpuVo `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizCode = ""
	m.Message = ""
	m.Data = nil
}

var poolTaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse)
	},
}

// GetTaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse 从 sync.Pool 获取 TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse
func GetTaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse() *TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse {
	return poolTaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse.Get().(*TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse)
}

// ReleaseTaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse 将 TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse(v *TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse) {
	v.Reset()
	poolTaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse.Put(v)
}
