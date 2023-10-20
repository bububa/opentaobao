package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodealerrequisitionorderremarkupdateAPIResponse 修改经销采购单备注 API返回值
// taobao.fenxiao.dealer.requisitionorder.remark.update
//
// 供应商修改经销采购单备注
type TaobaofenxiaodealerrequisitionorderremarkupdateAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaodealerrequisitionorderremarkupdateAPIResponseModel
}

// TaobaofenxiaodealerrequisitionorderremarkupdateAPIResponseModel is 修改经销采购单备注 成功返回结果
type TaobaofenxiaodealerrequisitionorderremarkupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_remark_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
