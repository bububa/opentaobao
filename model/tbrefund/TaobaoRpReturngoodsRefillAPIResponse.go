package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorpreturngoodsrefillAPIResponse 卖家回填物流信息 API返回值
// taobao.rp.returngoods.refill
//
// 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaorpreturngoodsrefillAPIResponse struct {
	model.CommonResponse
	TaobaorpreturngoodsrefillAPIResponseModel
}

// TaobaorpreturngoodsrefillAPIResponseModel is 卖家回填物流信息 成功返回结果
type TaobaorpreturngoodsrefillAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_returngoods_refill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 验货操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
