package usergrowth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘线下拉新业务 t+8转化明细数据 API返回值 
taobao.usergrowth.offline.convertion.details.eight.get

手淘线下拉新业务 给合作渠道返回t+8转化明细数据
*/
type TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIResponseModel
}

// 手淘线下拉新业务 t+8转化明细数据 成功返回结果
type TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIResponseModel struct {
    XMLName xml.Name `xml:"usergrowth_offline_convertion_details_eight_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 鹰眼id
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 返回数据的总条数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 每页多少条记录
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 页码
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 返回数据集合
    List   []TaobaoUsergrowthOfflineConvertionDetailsEightGetE `json:"list,omitempty" xml:"list>taobao_usergrowth_offline_convertion_details_eight_get_e,omitempty"`
}
