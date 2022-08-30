package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxAccountAccountFreezeAPIResponse 创建计划后支付 API返回值
// taobao.onebp.dkx.account.account.freeze
//
// 创建计划后支付。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxAccountAccountFreezeAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxAccountAccountFreezeAPIResponseModel
}

// TaobaoOnebpDkxAccountAccountFreezeAPIResponseModel is 创建计划后支付 成功返回结果
type TaobaoOnebpDkxAccountAccountFreezeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_account_account_freeze_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxAccountAccountFreezeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
