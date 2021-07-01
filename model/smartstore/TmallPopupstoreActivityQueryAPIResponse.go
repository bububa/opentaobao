package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某段时间内的快闪活动列表 API返回值 
tmall.popupstore.activity.query

提供给ISV查询某一时间段内包含指定appKey的活动列表
*/
type TmallPopupstoreActivityQueryAPIResponse struct {
    model.CommonResponse
    TmallPopupstoreActivityQueryAPIResponseModel
}

// 查询某段时间内的快闪活动列表 成功返回结果
type TmallPopupstoreActivityQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_popupstore_activity_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参对象
    ResultDto   *TmallPopupstoreActivityQueryResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}
