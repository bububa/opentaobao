package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaserchcrowdpricebatchupdateAPIResponse 单品推广搜索人群修改溢价 API返回值
// taobao.simba.serchcrowd.price.batch.update
//
// 单品推广搜索人群修改溢价, 不支持跨推广单元修改
type TaobaosimbaserchcrowdpricebatchupdateAPIResponse struct {
	model.CommonResponse
	TaobaosimbaserchcrowdpricebatchupdateAPIResponseModel
}

// TaobaosimbaserchcrowdpricebatchupdateAPIResponseModel is 单品推广搜索人群修改溢价 成功返回结果
type TaobaosimbaserchcrowdpricebatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_serchcrowd_price_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Adgrouptargetingtags []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty" xml:"adgrouptargetingtags>adgroup_targeting_tag_dto,omitempty"`
}
