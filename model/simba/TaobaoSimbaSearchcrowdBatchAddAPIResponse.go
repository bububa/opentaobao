package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasearchcrowdbatchaddAPIResponse 推广单元增加搜索人群 API返回值
// taobao.simba.searchcrowd.batch.add
//
// 推广单元新增搜索人群
type TaobaosimbasearchcrowdbatchaddAPIResponse struct {
	model.CommonResponse
	TaobaosimbasearchcrowdbatchaddAPIResponseModel
}

// TaobaosimbasearchcrowdbatchaddAPIResponseModel is 推广单元增加搜索人群 成功返回结果
type TaobaosimbasearchcrowdbatchaddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_searchcrowd_batch_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 定向信息
	Adgrouptargetingtags []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty" xml:"adgrouptargetingtags>adgroup_targeting_tag_dto,omitempty"`
}
