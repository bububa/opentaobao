package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateDelistingTmallAPIResponse taobao.item.update.delisting.tmall API返回值
// taobao.item.update.delisting.tmall
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingTmallAPIResponse struct {
	model.CommonResponse
	TaobaoItemUpdateDelistingTmallAPIResponseModel
}

// TaobaoItemUpdateDelistingTmallAPIResponseModel is taobao.item.update.delisting.tmall 成功返回结果
type TaobaoItemUpdateDelistingTmallAPIResponseModel struct {
	XMLName xml.Name `xml:"item_update_delisting_tmall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回商品更新信息：返回的结果是:num_iid和modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}
