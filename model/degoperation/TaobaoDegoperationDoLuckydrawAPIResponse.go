package degoperation

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationDoLuckydrawAPIResponse
激励抽奖 API返回值
taobao.degoperation.do.luckydraw

激励平台抽奖接口。用户可以通过接口完成抽奖功能 */
type TaobaoDegoperationDoLuckydrawAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationDoLuckydrawAPIResponseModel
}

// TaobaoDegoperationDoLuckydrawAPIResponseModel is 激励抽奖 成功返回结果
type TaobaoDegoperationDoLuckydrawAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_do_luckydraw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
