package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaochainstorerangesAPIResponse 蜂鸟查询门店配送范围接口 API返回值
// alibaba.ele.fengniao.chainstore.ranges
//
// 蜂鸟查询门店配送范围接口
type AlibabaelefengniaochainstorerangesAPIResponse struct {
	model.CommonResponse
	AlibabaelefengniaochainstorerangesAPIResponseModel
}

// AlibabaelefengniaochainstorerangesAPIResponseModel is 蜂鸟查询门店配送范围接口 成功返回结果
type AlibabaelefengniaochainstorerangesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_ranges_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	RangeList []AlibabaelefengniaochainstorerangesResult `json:"range_list,omitempty" xml:"range_list>alibabaelefengniaochainstoreranges_result,omitempty"`
}
