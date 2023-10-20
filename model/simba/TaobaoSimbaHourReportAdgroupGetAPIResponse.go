package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbahourreportadgroupgetAPIResponse 推广单元小时级别实时报表查询 API返回值
// taobao.simba.hour.report.adgroup.get
//
// 推广单元小时级别实时报表查询
type TaobaosimbahourreportadgroupgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbahourreportadgroupgetAPIResponseModel
}

// TaobaosimbahourreportadgroupgetAPIResponseModel is 推广单元小时级别实时报表查询 成功返回结果
type TaobaosimbahourreportadgroupgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_hour_report_adgroup_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 11
	Results *RtRptResultEntityDto `json:"results,omitempty" xml:"results,omitempty"`
}
