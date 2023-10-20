package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartapptablefulldatagetAPIResponse 智能应用工作表地址查询 API返回值
// taobao.smartapp.table.fulldata.get
//
// 智能应用工作表地址查询
type TaobaosmartapptablefulldatagetAPIResponse struct {
	model.CommonResponse
	TaobaosmartapptablefulldatagetAPIResponseModel
}

// TaobaosmartapptablefulldatagetAPIResponseModel is 智能应用工作表地址查询 成功返回结果
type TaobaosmartapptablefulldatagetAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_fulldata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 查询结果
	Result *AfterSaleTableSelectResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}
