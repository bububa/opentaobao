package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse 推广组下创意报表基础数据查询(汇总数据，不分类型) API返回值
// taobao.simba.rpt.adgroupcreativebase.get
//
// 推广组下创意报表基础数据查询(汇总数据，不分类型)
type TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel
}

// TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel is 推广组下创意报表基础数据查询(汇总数据，不分类型) 成功返回结果
type TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupcreativebase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组下的创意基础数据列表
	RptAdgroupcreativeBaseList string `json:"rpt_adgroupcreative_base_list,omitempty" xml:"rpt_adgroupcreative_base_list,omitempty"`
}
