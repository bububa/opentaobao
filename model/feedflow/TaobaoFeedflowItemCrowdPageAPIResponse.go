package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdPageAPIResponse
分页查询单品单元下人群列表 API返回值
taobao.feedflow.item.crowd.page

分页查询单品单元下人群列表 */
type TaobaoFeedflowItemCrowdPageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCrowdPageAPIResponseModel
}

// TaobaoFeedflowItemCrowdPageAPIResponseModel is 分页查询单品单元下人群列表 成功返回结果
type TaobaoFeedflowItemCrowdPageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemCrowdPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
