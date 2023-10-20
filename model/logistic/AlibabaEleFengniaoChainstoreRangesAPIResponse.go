package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoChainstoreRangesAPIResponse 蜂鸟查询门店配送范围接口 API返回值
// alibaba.ele.fengniao.chainstore.ranges
//
// 蜂鸟查询门店配送范围接口
type AlibabaEleFengniaoChainstoreRangesAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoChainstoreRangesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoChainstoreRangesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoChainstoreRangesAPIResponseModel).Reset()
}

// AlibabaEleFengniaoChainstoreRangesAPIResponseModel is 蜂鸟查询门店配送范围接口 成功返回结果
type AlibabaEleFengniaoChainstoreRangesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_ranges_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	RangeList []AlibabaEleFengniaoChainstoreRangesResult `json:"range_list,omitempty" xml:"range_list>alibaba_ele_fengniao_chainstore_ranges_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoChainstoreRangesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RangeList = m.RangeList[:0]
}

var poolAlibabaEleFengniaoChainstoreRangesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoChainstoreRangesAPIResponse)
	},
}

// GetAlibabaEleFengniaoChainstoreRangesAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoChainstoreRangesAPIResponse
func GetAlibabaEleFengniaoChainstoreRangesAPIResponse() *AlibabaEleFengniaoChainstoreRangesAPIResponse {
	return poolAlibabaEleFengniaoChainstoreRangesAPIResponse.Get().(*AlibabaEleFengniaoChainstoreRangesAPIResponse)
}

// ReleaseAlibabaEleFengniaoChainstoreRangesAPIResponse 将 AlibabaEleFengniaoChainstoreRangesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoChainstoreRangesAPIResponse(v *AlibabaEleFengniaoChainstoreRangesAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoChainstoreRangesAPIResponse.Put(v)
}
