package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaChongzhiQueryecardsAPIResponse 查询指定商家的可用的话费宝贝接口 API返回值
// alibaba.chongzhi.queryecards
//
// 查询指定商家的可用的话费宝贝
type AlibabaChongzhiQueryecardsAPIResponse struct {
	model.CommonResponse
	AlibabaChongzhiQueryecardsAPIResponseModel
}

// AlibabaChongzhiQueryecardsAPIResponseModel is 查询指定商家的可用的话费宝贝接口 成功返回结果
type AlibabaChongzhiQueryecardsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_chongzhi_queryecards_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单列表
	EcardList []EcardItemDo `json:"ecard_list,omitempty" xml:"ecard_list>ecard_item_do,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// MtsInfoDo
	CatInfo *MtsInfoDo `json:"cat_info,omitempty" xml:"cat_info,omitempty"`
}
