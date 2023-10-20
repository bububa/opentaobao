package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableUpdateAPIResponse 智能应用服务登记工作表数据更新 API返回值
// taobao.smartapp.table.update
//
// 智能应用服务登记工作表数据更新
type TaobaoSmartappTableUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappTableUpdateAPIResponseModel
}

// TaobaoSmartappTableUpdateAPIResponseModel is 智能应用服务登记工作表数据更新 成功返回结果
type TaobaoSmartappTableUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 更新记录行数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}
