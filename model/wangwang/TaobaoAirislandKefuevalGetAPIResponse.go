package wangwang

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAirislandKefuevalGetAPIResponse 客服评价详情接口_V2 API返回值
// taobao.airisland.kefueval.get
//
// 获取买家对客服的服务评价
//
// 注意：
//
// 1. 请求超时[isp.top-remote-connection-timeout]或者数据过大错误[isp.runtime-max-limit]：因为某些帐号请求的数据会非常大，【需要通过减少请求时间范围避免该问题】
//
// 2. 时间范围：[now()-90d&lt;=btime ~ etime &lt;= now()-1d ] AND etime-btime &lt;=7d
//
// 3. 变更eval_recer：可空，返回脱敏的买家nick，如：摩天轮 -&gt; 摩**
//
// 4. 新增labelName：可空
type TaobaoAirislandKefuevalGetAPIResponse struct {
	model.CommonResponse
	TaobaoAirislandKefuevalGetAPIResponseModel
}

// TaobaoAirislandKefuevalGetAPIResponseModel is 客服评价详情接口_V2 成功返回结果
type TaobaoAirislandKefuevalGetAPIResponseModel struct {
	XMLName xml.Name `xml:"airisland_kefueval_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 评价明细
	StaffEvalDetails []EvalDetail `json:"staff_eval_details,omitempty" xml:"staff_eval_details>eval_detail,omitempty"`
	// 评价结果数
	ResultCount int64 `json:"result_count,omitempty" xml:"result_count,omitempty"`
}
