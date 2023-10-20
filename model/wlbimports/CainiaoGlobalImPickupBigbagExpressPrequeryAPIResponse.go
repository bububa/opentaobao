package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupbigbagexpressprequeryAPIResponse 首公里揽收-快递预查询服务 API返回值
// cainiao.global.im.pickup.bigbag.express.prequery
//
// 快递预查询服务
type CainiaoglobalimpickupbigbagexpressprequeryAPIResponse struct {
	model.CommonResponse
	CainiaoglobalimpickupbigbagexpressprequeryAPIResponseModel
}

// CainiaoglobalimpickupbigbagexpressprequeryAPIResponseModel is 首公里揽收-快递预查询服务 成功返回结果
type CainiaoglobalimpickupbigbagexpressprequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_express_prequery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
