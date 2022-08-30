package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthOfflineNewloginQueryAPIResponse 线下新登数据实时查询接口 API返回值
// taobao.usergrowth.offline.newlogin.query
//
// 线下新登数据实时查询接口
type TaobaoUsergrowthOfflineNewloginQueryAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthOfflineNewloginQueryAPIResponseModel
}

// TaobaoUsergrowthOfflineNewloginQueryAPIResponseModel is 线下新登数据实时查询接口 成功返回结果
type TaobaoUsergrowthOfflineNewloginQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_offline_newlogin_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功获取数据
	Successful string `json:"successful,omitempty" xml:"successful,omitempty"`
	// 新登数据
	Data *OfflinePromoteInfo `json:"data,omitempty" xml:"data,omitempty"`
}
