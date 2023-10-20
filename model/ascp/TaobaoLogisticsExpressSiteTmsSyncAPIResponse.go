package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresssitetmssyncAPIResponse 配服务商网点信息同步 API返回值
// taobao.logistics.express.site.tms.sync
//
// 配服务商网点信息同步
type TaobaologisticsexpresssitetmssyncAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpresssitetmssyncAPIResponseModel
}

// TaobaologisticsexpresssitetmssyncAPIResponseModel is 配服务商网点信息同步 成功返回结果
type TaobaologisticsexpresssitetmssyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_site_tms_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	SiteUpsetResponse *SiteUpsetResponse `json:"site_upset_response,omitempty" xml:"site_upset_response,omitempty"`
}
