package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemitempublishAPIResponse 全渠道门店商品轻发布 API返回值
// taobao.omniitem.item.publish
//
// 全渠道门店商品轻发布
type TaobaoomniitemitempublishAPIResponse struct {
	model.CommonResponse
	TaobaoomniitemitempublishAPIResponseModel
}

// TaobaoomniitemitempublishAPIResponseModel is 全渠道门店商品轻发布 成功返回结果
type TaobaoomniitemitempublishAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_item_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ItemLightPublishResult `json:"result,omitempty" xml:"result,omitempty"`
}
