package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAdvertQueryscheduleAPIResponse 广告牌照管控查询 API返回值
// yunos.tvpubadmin.content.advert.queryschedule
//
// 广告牌照管控查询
type YunosTvpubadminContentAdvertQueryscheduleAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentAdvertQueryscheduleAPIResponseModel
}

// YunosTvpubadminContentAdvertQueryscheduleAPIResponseModel is 广告牌照管控查询 成功返回结果
type YunosTvpubadminContentAdvertQueryscheduleAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_queryschedule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 具体数据信息
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}
