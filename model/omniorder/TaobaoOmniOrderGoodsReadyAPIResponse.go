package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniordergoodsreadyAPIResponse 备货完成 API返回值
// taobao.omni.order.goods.ready
//
// 备货完成
type TaobaoomniordergoodsreadyAPIResponse struct {
	model.CommonResponse
	TaobaoomniordergoodsreadyAPIResponseModel
}

// TaobaoomniordergoodsreadyAPIResponseModel is 备货完成 成功返回结果
type TaobaoomniordergoodsreadyAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_order_goods_ready_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
