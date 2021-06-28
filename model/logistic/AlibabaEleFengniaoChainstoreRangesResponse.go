package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟查询门店配送范围接口 APIResponse
alibaba.ele.fengniao.chainstore.ranges

蜂鸟查询门店配送范围接口
*/
type AlibabaEleFengniaoChainstoreRangesAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaEleFengniaoChainstoreRangesResponse `json:"alibaba_ele_fengniao_chainstore_ranges_response,omitempty"` 
    AlibabaEleFengniaoChainstoreRangesResponse
}

/* model for simplify = false
type AlibabaEleFengniaoChainstoreRangesResponse struct {

    // 返回值
    
    RangeList  struct {
        AlibabaEleFengniaoChainstoreRangesResult  []AlibabaEleFengniaoChainstoreRangesResult `json:"alibaba_ele_fengniao_chainstore_ranges_result,omitempty"`
    } `json:"range_list,omitempty"`
    

}
*/

type AlibabaEleFengniaoChainstoreRangesResponse struct {

    // 返回值
    RangeList   []AlibabaEleFengniaoChainstoreRangesResult `json:"range_list,omitempty"`

}
