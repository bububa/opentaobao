package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadtargettaggetallenabletaglistAPIResponse 查询所有可添加标签信息 API返回值
// alibaba.scbp.ad.target.tag.get.all.enable.tag.list
//
// 查询标签数据
type AlibabascbpadtargettaggetallenabletaglistAPIResponse struct {
	model.CommonResponse
	AlibabascbpadtargettaggetallenabletaglistAPIResponseModel
}

// AlibabascbpadtargettaggetallenabletaglistAPIResponseModel is 查询所有可添加标签信息 成功返回结果
type AlibabascbpadtargettaggetallenabletaglistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_get_all_enable_tag_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回实体
	ResultList []AdsTargetingTagDto `json:"result_list,omitempty" xml:"result_list>ads_targeting_tag_dto,omitempty"`
}
