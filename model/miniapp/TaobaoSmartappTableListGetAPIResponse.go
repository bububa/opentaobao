package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartapptablelistgetAPIResponse 智能应用服务登记工作表列表查询 API返回值
// taobao.smartapp.table.list.get
//
// 智能应用服务登记工作表列表查询
type TaobaosmartapptablelistgetAPIResponse struct {
	model.CommonResponse
	TaobaosmartapptablelistgetAPIResponseModel
}

// TaobaosmartapptablelistgetAPIResponseModel is 智能应用服务登记工作表列表查询 成功返回结果
type TaobaosmartapptablelistgetAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用结果
	Result *AfterSaleGetWorkTableListResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}
