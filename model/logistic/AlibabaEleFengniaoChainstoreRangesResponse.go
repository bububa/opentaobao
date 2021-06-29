package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟查询门店配送范围接口 API返回值 
alibaba.ele.fengniao.chainstore.ranges

蜂鸟查询门店配送范围接口
*/
type AlibabaEleFengniaoChainstoreRangesAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoChainstoreRangesResponse
}

// 蜂鸟查询门店配送范围接口 成功返回结果
type AlibabaEleFengniaoChainstoreRangesResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_ranges_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    RangeList   []AlibabaEleFengniaoChainstoreRangesResult `json:"range_list,omitempty" xml:"range_list>alibaba_ele_fengniao_chainstore_ranges_result,omitempty"`
}
