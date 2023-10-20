package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcrowdfindrecommendcrowdAPIResponse 查询推荐人群 API返回值
// taobao.universalbp.crowd.findrecommendcrowd
//
// 入参推广信息，查询推荐人群，查出的推荐人群列表，不需转换，直接可用于入参其他计划创编类接口
type TaobaouniversalbpcrowdfindrecommendcrowdAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpcrowdfindrecommendcrowdAPIResponseModel
}

// TaobaouniversalbpcrowdfindrecommendcrowdAPIResponseModel is 查询推荐人群 成功返回结果
type TaobaouniversalbpcrowdfindrecommendcrowdAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_crowd_findrecommendcrowd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpcrowdfindrecommendcrowdTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
