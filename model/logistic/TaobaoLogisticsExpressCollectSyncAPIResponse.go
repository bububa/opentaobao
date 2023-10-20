package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresscollectsyncAPIResponse 服饰逆向揽收信息同步 API返回值
// taobao.logistics.express.collect.sync
//
// 服饰逆向揽收信息同步
type TaobaologisticsexpresscollectsyncAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpresscollectsyncAPIResponseModel
}

// TaobaologisticsexpresscollectsyncAPIResponseModel is 服饰逆向揽收信息同步 成功返回结果
type TaobaologisticsexpresscollectsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_collect_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码描述
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码标识
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 返回值
	Data *TmsCollectResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 校验成功或者异常
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
