package tmallhk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearancegetAPIResponse 天猫国际-清关材料查询 API返回值
// tmall.hk.clearance.get
//
// 提供订单收货人身份信息查询功能。
type TmallhkclearancegetAPIResponse struct {
	model.CommonResponse
	TmallhkclearancegetAPIResponseModel
}

// TmallhkclearancegetAPIResponseModel is 天猫国际-清关材料查询 成功返回结果
type TmallhkclearancegetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_hk_clearance_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果对象
	Result *CertifyQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
