package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoOrderRemarkUpdateAPIResponse 修改采购单备注 API返回值
// taobao.fenxiao.order.remark.update
//
// 供应商修改采购单备注
type TaobaoFenxiaoOrderRemarkUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel
}

// TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel is 修改采购单备注 成功返回结果
type TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_order_remark_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
