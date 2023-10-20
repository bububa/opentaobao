package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpdetailupdateAPIResponse 修改活动详情 API返回值
// taobao.ump.detail.update
//
// 更新活动详情
type TaobaoumpdetailupdateAPIResponse struct {
	model.CommonResponse
	TaobaoumpdetailupdateAPIResponseModel
}

// TaobaoumpdetailupdateAPIResponseModel is 修改活动详情 成功返回结果
type TaobaoumpdetailupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_detail_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
