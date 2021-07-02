package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStallPayratioSynchronizeAPIResponse 同步摊位收银比例 API返回值
// tmall.nrt.stall.payratio.synchronize
//
// ISV同步摊位收银比例到阿里
type TmallNrtStallPayratioSynchronizeAPIResponse struct {
	model.CommonResponse
	TmallNrtStallPayratioSynchronizeAPIResponseModel
}

// TmallNrtStallPayratioSynchronizeAPIResponseModel is 同步摊位收银比例 成功返回结果
type TmallNrtStallPayratioSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_stall_payratio_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
