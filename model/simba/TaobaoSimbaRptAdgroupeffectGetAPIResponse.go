package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptAdgroupeffectGetAPIResponse
推广组效果报表数据对象 API返回值
taobao.simba.rpt.adgroupeffect.get

推广组效果报表数据对象 */
type TaobaoSimbaRptAdgroupeffectGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptAdgroupeffectGetAPIResponseModel
}

// TaobaoSimbaRptAdgroupeffectGetAPIResponseModel is 推广组效果报表数据对象 成功返回结果
type TaobaoSimbaRptAdgroupeffectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupeffect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组效果报表数据对象
	RptAdgroupEffectList string `json:"rpt_adgroup_effect_list,omitempty" xml:"rpt_adgroup_effect_list,omitempty"`
}
