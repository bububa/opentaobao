package qt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqtreportaddAPIResponse 上传质检报告 API返回值
// taobao.qt.report.add
//
// 上传质检报告
type TaobaoqtreportaddAPIResponse struct {
	model.CommonResponse
	TaobaoqtreportaddAPIResponseModel
}

// TaobaoqtreportaddAPIResponseModel is 上传质检报告 成功返回结果
type TaobaoqtreportaddAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_report_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
