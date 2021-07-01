package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmMemberScoreGetAPIResponse
获取会员某时间段积分 API返回值
alibaba.westcrm.member.score.get

获取会员某时间段积分 */
type AlibabaWestcrmMemberScoreGetAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmMemberScoreGetAPIResponseModel
}

// AlibabaWestcrmMemberScoreGetAPIResponseModel is 获取会员某时间段积分 成功返回结果
type AlibabaWestcrmMemberScoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_member_score_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
