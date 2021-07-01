package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询定向标签列表 API返回值 
taobao.feedflow.item.option.page

分页查询定向标签列表
*/
type TaobaoFeedflowItemOptionPageAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemOptionPageAPIResponseModel
}

// 分页查询定向标签列表 成功返回结果
type TaobaoFeedflowItemOptionPageAPIResponseModel struct {
    XMLName xml.Name `xml:"feedflow_item_option_page_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *TaobaoFeedflowItemOptionPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
