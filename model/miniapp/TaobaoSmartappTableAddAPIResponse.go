package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartapptableaddAPIResponse 智能应用服务登记工作表数据新增 API返回值
// taobao.smartapp.table.add
//
// 智能应用服务登记工作表数据新增
type TaobaosmartapptableaddAPIResponse struct {
	model.CommonResponse
	TaobaosmartapptableaddAPIResponseModel
}

// TaobaosmartapptableaddAPIResponseModel is 智能应用服务登记工作表数据新增 成功返回结果
type TaobaosmartapptableaddAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单条新增记录ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 错误提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}
