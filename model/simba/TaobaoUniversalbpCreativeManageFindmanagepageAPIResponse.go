package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse 创意库查询创意列表 API返回值
// taobao.universalbp.creative.manage.findmanagepage
//
// 创意库查询创意列表
type TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel
}

// TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel is 创意库查询创意列表 成功返回结果
type TaobaoUniversalbpCreativeManageFindmanagepageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_creative_manage_findmanagepage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCreativeManageFindmanagepageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
