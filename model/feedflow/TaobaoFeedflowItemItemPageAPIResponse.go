package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemItemPageAPIResponse 信息流查看商品列表 API返回值
// taobao.feedflow.item.item.page
//
// 信息流查看商品列表
type TaobaoFeedflowItemItemPageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemItemPageAPIResponseModel
}

// TaobaoFeedflowItemItemPageAPIResponseModel is 信息流查看商品列表 成功返回结果
type TaobaoFeedflowItemItemPageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_item_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemItemPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
