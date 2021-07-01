package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPaimaiItempropsGetAPIResponse
拍卖相关类目属性 API返回值
taobao.paimai.itemprops.get

读取拍卖相关类目属性 */
type TaobaoPaimaiItempropsGetAPIResponse struct {
	model.CommonResponse
	TaobaoPaimaiItempropsGetAPIResponseModel
}

// TaobaoPaimaiItempropsGetAPIResponseModel is 拍卖相关类目属性 成功返回结果
type TaobaoPaimaiItempropsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_itemprops_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果
	ItemProps []ItemProp `json:"item_props,omitempty" xml:"item_props>item_prop,omitempty"`
	// lastModified
	LastModified string `json:"last_modified,omitempty" xml:"last_modified,omitempty"`
}
