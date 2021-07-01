package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemUpdateDelistingAPIResponse
商品下架 API返回值
taobao.item.update.delisting

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateDelistingAPIResponse struct {
	model.CommonResponse
	TaobaoItemUpdateDelistingAPIResponseModel
}

// TaobaoItemUpdateDelistingAPIResponseModel is 商品下架 成功返回结果
type TaobaoItemUpdateDelistingAPIResponseModel struct {
	XMLName xml.Name `xml:"item_update_delisting_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回商品更新信息：返回的结果是:num_iid和modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}
