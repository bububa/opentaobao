package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopipoutgetAPIResponse 获取开放平台出口IP段 API返回值
// taobao.top.ipout.get
//
// 获取开放平台出口IP段
type TaobaotopipoutgetAPIResponse struct {
	model.CommonResponse
	TaobaotopipoutgetAPIResponseModel
}

// TaobaotopipoutgetAPIResponseModel is 获取开放平台出口IP段 成功返回结果
type TaobaotopipoutgetAPIResponseModel struct {
	XMLName xml.Name `xml:"top_ipout_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TOP网关出口IP列表
	IpList string `json:"ip_list,omitempty" xml:"ip_list,omitempty"`
}
