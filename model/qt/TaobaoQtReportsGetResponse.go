package qt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询质检报告 APIResponse
taobao.qt.reports.get

批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
*/
type TaobaoQtReportsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qt_reports_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 质检报告列表
    
    Reports   []QtReport `json:"reports,omitempty" xml:"