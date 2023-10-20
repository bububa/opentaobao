package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaovirtualdsfsupplierinterfaceswitchAPIResponse 虚拟供应商履约接口切换 API返回值
// taobao.virtual.dsf.supplier.interface.switch
//
// 虚拟供应商履约接口切换
type TaobaovirtualdsfsupplierinterfaceswitchAPIResponse struct {
	model.CommonResponse
	TaobaovirtualdsfsupplierinterfaceswitchAPIResponseModel
}

// TaobaovirtualdsfsupplierinterfaceswitchAPIResponseModel is 虚拟供应商履约接口切换 成功返回结果
type TaobaovirtualdsfsupplierinterfaceswitchAPIResponseModel struct {
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
