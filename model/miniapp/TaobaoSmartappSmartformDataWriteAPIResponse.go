package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappSmartformDataWriteAPIResponse 智能表单外部更新数据 API返回值
// taobao.smartapp.smartform.data.write
//
// 智能表单外部更新数据
type TaobaoSmartappSmartformDataWriteAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappSmartformDataWriteAPIResponseModel
}

// TaobaoSmartappSmartformDataWriteAPIResponseModel is 智能表单外部更新数据 成功返回结果
type TaobaoSmartappSmartformDataWriteAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_smartform_data_write_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作数据ID
	RecordId string `json:"record_id,omitempty" xml:"record_id,omitempty"`
}
