package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimensuppliersynchronizeAPIResponse 供应商同步接口 API返回值
// taobao.qimen.supplier.synchronize
//
// 这个接口用来同步供应商信息
type TaobaoqimensuppliersynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoqimensuppliersynchronizeAPIResponseModel
}

// TaobaoqimensuppliersynchronizeAPIResponseModel is 供应商同步接口 成功返回结果
type TaobaoqimensuppliersynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_supplier_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimensuppliersynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
