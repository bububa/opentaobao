package qt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqtreportupdateAPIResponse 更新质检报告 API返回值
// taobao.qt.report.update
//
// 更新质检报告
type TaobaoqtreportupdateAPIResponse struct {
	model.CommonResponse
	TaobaoqtreportupdateAPIResponseModel
}

// TaobaoqtreportupdateAPIResponseModel is 更新质检报告 成功返回结果
type TaobaoqtreportupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_report_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
