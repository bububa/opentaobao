package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupAdzonePageAPIResponse
信息流单元下查看绑定资源位 API返回值
taobao.feedflow.item.adgroup.adzone.page

信息流单元下查看绑定资源位 */
type TaobaoFeedflowItemAdgroupAdzonePageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel
}

// TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel is 信息流单元下查看绑定资源位 成功返回结果
type TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAdgroupAdzonePageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
