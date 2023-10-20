package koubeimall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonitemsuperdiscountlistAPIResponse 查询商圈内的超值特惠商品信息 API返回值
// taobao.koubei.mall.common.item.super.discount.list
//
// 查询商圈超值特惠商品信息列表
type TaobaokoubeimallcommonitemsuperdiscountlistAPIResponse struct {
	model.CommonResponse
	TaobaokoubeimallcommonitemsuperdiscountlistAPIResponseModel
}

// TaobaokoubeimallcommonitemsuperdiscountlistAPIResponseModel is 查询商圈内的超值特惠商品信息 成功返回结果
type TaobaokoubeimallcommonitemsuperdiscountlistAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_item_super_discount_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaokoubeimallcommonitemsuperdiscountlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
