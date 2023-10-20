package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodealerrequisitionordercreateAPIResponse 创建经销采购申请 API返回值
// taobao.fenxiao.dealer.requisitionorder.create
//
// 创建经销采购申请
type TaobaofenxiaodealerrequisitionordercreateAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaodealerrequisitionordercreateAPIResponseModel
}

// TaobaofenxiaodealerrequisitionordercreateAPIResponseModel is 创建经销采购申请 成功返回结果
type TaobaofenxiaodealerrequisitionordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 经销采购申请编号
	DealerOrderId int64 `json:"dealer_order_id,omitempty" xml:"dealer_order_id,omitempty"`
}
