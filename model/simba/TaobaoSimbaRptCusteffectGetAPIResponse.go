package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptCusteffectGetAPIResponse
用户账户报表效果数据查询（只有汇总数据，无分类数据） API返回值
taobao.simba.rpt.custeffect.get

用户账户报表效果数据查询（只有汇总数据，无分类数据） */
type TaobaoSimbaRptCusteffectGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptCusteffectGetAPIResponseModel
}

// TaobaoSimbaRptCusteffectGetAPIResponseModel is 用户账户报表效果数据查询（只有汇总数据，无分类数据） 成功返回结果
type TaobaoSimbaRptCusteffectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_custeffect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账户效果数据返回结果
	RptCustEffectList string `json:"rpt_cust_effect_list,omitempty" xml:"rpt_cust_effect_list,omitempty"`
}
