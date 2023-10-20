package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorpreturngoodsagreeAPIResponse 卖家同意退货 API返回值
// taobao.rp.returngoods.agree
//
// 卖家同意退货，支持淘宝和天猫的订单。
type TaobaorpreturngoodsagreeAPIResponse struct {
	model.CommonResponse
	TaobaorpreturngoodsagreeAPIResponseModel
}

// TaobaorpreturngoodsagreeAPIResponseModel is 卖家同意退货 成功返回结果
type TaobaorpreturngoodsagreeAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_returngoods_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
