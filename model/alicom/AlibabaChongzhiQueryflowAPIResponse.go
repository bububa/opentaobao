package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaChongzhiQueryflowAPIResponse
查询指定商家的可用的流量宝贝接口 API返回值
alibaba.chongzhi.queryflow

查询指定商家的可用的流量宝贝 */
type AlibabaChongzhiQueryflowAPIResponse struct {
	model.CommonResponse
	AlibabaChongzhiQueryflowAPIResponseModel
}

// AlibabaChongzhiQueryflowAPIResponseModel is 查询指定商家的可用的流量宝贝接口 成功返回结果
type AlibabaChongzhiQueryflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_chongzhi_queryflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// MtsInfoDo
	CatInfo *MtsInfoDo `json:"cat_info,omitempty" xml:"cat_info,omitempty"`
	// flow_card_list
	FlowCardList []Flowcardlist `json:"flow_card_list,omitempty" xml:"flow_card_list>flowcardlist,omitempty"`
	// cn_desc
	CnDesc string `json:"cn_desc,omitempty" xml:"cn_desc,omitempty"`
}
