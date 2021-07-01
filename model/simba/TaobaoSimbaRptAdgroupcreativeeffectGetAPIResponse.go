package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse
推广组下的创意报表效果数据查询(汇总数据，不分类型) API返回值
taobao.simba.rpt.adgroupcreativeeffect.get

推广组下的创意报表效果数据查询(汇总数据，不分类型) */
type TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel
}

// TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel is 推广组下的创意报表效果数据查询(汇总数据，不分类型) 成功返回结果
type TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupcreativeeffect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组下的创意效果数据列表
	RptAdgroupcreativeEffectList string `json:"rpt_adgroupcreative_effect_list,omitempty" xml:"rpt_adgroupcreative_effect_list,omitempty"`
}
