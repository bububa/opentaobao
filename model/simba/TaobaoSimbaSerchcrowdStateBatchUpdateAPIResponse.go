package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaserchcrowdstatebatchupdateAPIResponse 单品搜索人群修改状态 API返回值
// taobao.simba.serchcrowd.state.batch.update
//
// 暂停或启用单品推广搜索人群溢价
type TaobaosimbaserchcrowdstatebatchupdateAPIResponse struct {
	model.CommonResponse
	TaobaosimbaserchcrowdstatebatchupdateAPIResponseModel
}

// TaobaosimbaserchcrowdstatebatchupdateAPIResponseModel is 单品搜索人群修改状态 成功返回结果
type TaobaosimbaserchcrowdstatebatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_serchcrowd_state_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Adgrouptargetingtags []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty" xml:"adgrouptargetingtags>adgroup_targeting_tag_dto,omitempty"`
	// 部分失败时返回错误List
	ErrorList []string `json:"error_list,omitempty" xml:"error_list>string,omitempty"`
}
