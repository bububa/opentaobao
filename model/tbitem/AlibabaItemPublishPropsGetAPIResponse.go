package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitempublishpropsgetAPIResponse 商品级联属性信息获取 API返回值
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
type AlibabaitempublishpropsgetAPIResponse struct {
	model.CommonResponse
	AlibabaitempublishpropsgetAPIResponseModel
}

// AlibabaitempublishpropsgetAPIResponseModel is 商品级联属性信息获取 成功返回结果
type AlibabaitempublishpropsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_publish_props_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品发布规则信息，XML格式.
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
