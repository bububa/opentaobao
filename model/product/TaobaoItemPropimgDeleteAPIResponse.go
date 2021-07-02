package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPropimgDeleteAPIResponse 删除属性图片 API返回值
// taobao.item.propimg.delete
//
// 删除propimg_id 所指定的商品属性图片 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>propimg_id对应的属性图片需要属于num_iid对应的商品
type TaobaoItemPropimgDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoItemPropimgDeleteAPIResponseModel
}

// TaobaoItemPropimgDeleteAPIResponseModel is 删除属性图片 成功返回结果
type TaobaoItemPropimgDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"item_propimg_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 属性图片结构
	PropImg *PropImg `json:"prop_img,omitempty" xml:"prop_img,omitempty"`
}
