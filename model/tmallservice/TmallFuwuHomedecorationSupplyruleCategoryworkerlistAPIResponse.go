package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIResponse 基于规则查品牌品类工人接口 API返回值
// tmall.fuwu.homedecoration.supplyrule.categoryworkerlist
//
// 基于规则查品牌品类工人接口
type TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIResponse struct {
	model.CommonResponse
	TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIResponseModel
}

// TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIResponseModel is 基于规则查品牌品类工人接口 成功返回结果
type TmallfuwuhomedecorationsupplyrulecategoryworkerlistAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_homedecoration_supplyrule_categoryworkerlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *TmallfuwuhomedecorationsupplyrulecategoryworkerlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
