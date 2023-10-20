package uscesl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizapaddAPIResponse 新增价签通讯AP设备 API返回值
// taobao.uscesl.biz.ap.add
//
// 根据门店和ap的MAC地址新增
type TaobaousceslbizapaddAPIResponse struct {
	model.CommonResponse
	TaobaousceslbizapaddAPIResponseModel
}

// TaobaousceslbizapaddAPIResponseModel is 新增价签通讯AP设备 成功返回结果
type TaobaousceslbizapaddAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_ap_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaousceslbizapaddResult `json:"result,omitempty" xml:"result,omitempty"`
}
