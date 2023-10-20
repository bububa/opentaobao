package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaserchcrowdgetAPIResponse 根据推广单元id获取搜索溢价人群 API返回值
// taobao.simba.serchcrowd.get
//
// 根据推广单元id获取搜索溢价人群
type TaobaosimbaserchcrowdgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbaserchcrowdgetAPIResponseModel
}

// TaobaosimbaserchcrowdgetAPIResponseModel is 根据推广单元id获取搜索溢价人群 成功返回结果
type TaobaosimbaserchcrowdgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_serchcrowd_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Adgrouptargetingtags []TaobaosimbaserchcrowdgetResult `json:"adgrouptargetingtags,omitempty" xml:"adgrouptargetingtags>taobaosimbaserchcrowdget_result,omitempty"`
}
