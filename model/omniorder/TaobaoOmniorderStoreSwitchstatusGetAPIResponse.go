package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoreswitchstatusgetAPIResponse switchstatus.get API返回值
// taobao.omniorder.store.switchstatus.get
//
// 查询门店发货、门店自提状态
type TaobaoomniorderstoreswitchstatusgetAPIResponse struct {
	model.CommonResponse
	TaobaoomniorderstoreswitchstatusgetAPIResponseModel
}

// TaobaoomniorderstoreswitchstatusgetAPIResponseModel is switchstatus.get 成功返回结果
type TaobaoomniorderstoreswitchstatusgetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_switchstatus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoomniorderstoreswitchstatusgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
