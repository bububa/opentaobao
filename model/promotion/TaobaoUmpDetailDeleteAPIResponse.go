package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpDetailDeleteAPIResponse
删除活动详情 API返回值
taobao.ump.detail.delete

删除活动详情 */
type TaobaoUmpDetailDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoUmpDetailDeleteAPIResponseModel
}

// TaobaoUmpDetailDeleteAPIResponseModel is 删除活动详情 成功返回结果
type TaobaoUmpDetailDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_detail_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
