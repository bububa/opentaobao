package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbareportcitygetAPIResponse 获取城市维度报表 API返回值
// taobao.simba.report.city.get
//
// 获取城市维度报表
type TaobaosimbareportcitygetAPIResponse struct {
	model.CommonResponse
	TaobaosimbareportcitygetAPIResponseModel
}

// TaobaosimbareportcitygetAPIResponseModel is 获取城市维度报表 成功返回结果
type TaobaosimbareportcitygetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_report_city_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 11
	Results *RtRptResultEntityDto `json:"results,omitempty" xml:"results,omitempty"`
}
