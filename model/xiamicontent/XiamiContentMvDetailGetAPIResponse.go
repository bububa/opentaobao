package xiamicontent

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentMvDetailGetAPIResponse
获取mv详情 API返回值
xiami.content.mv.detail.get

获取mv详情 */
type XiamiContentMvDetailGetAPIResponse struct {
	model.CommonResponse
	XiamiContentMvDetailGetAPIResponseModel
}

// XiamiContentMvDetailGetAPIResponseModel is 获取mv详情 成功返回结果
type XiamiContentMvDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_mv_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// mv列表
	MvDtos []MvDto `json:"mv_dtos,omitempty" xml:"mv_dtos>mv_dto,omitempty"`
	// 请求结果
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
