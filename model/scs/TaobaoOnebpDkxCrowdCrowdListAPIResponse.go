package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcrowdcrowdlistAPIResponse 获取人群信息列表 API返回值
// taobao.onebp.dkx.crowd.crowd.list
//
// 获取人群信息列表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxcrowdcrowdlistAPIResponse struct {
	model.CommonResponse
	TaobaoonebpdkxcrowdcrowdlistAPIResponseModel
}

// TaobaoonebpdkxcrowdcrowdlistAPIResponseModel is 获取人群信息列表 成功返回结果
type TaobaoonebpdkxcrowdcrowdlistAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_crowd_crowd_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoonebpdkxcrowdcrowdlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
