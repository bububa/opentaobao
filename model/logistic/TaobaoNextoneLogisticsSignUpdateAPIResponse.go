package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaonextonelogisticssignupdateAPIResponse AG物流签收状态写接口 API返回值
// taobao.nextone.logistics.sign.update
//
// 商家上传退货的签收状态给AG
type TaobaonextonelogisticssignupdateAPIResponse struct {
	model.CommonResponse
	TaobaonextonelogisticssignupdateAPIResponseModel
}

// TaobaonextonelogisticssignupdateAPIResponseModel is AG物流签收状态写接口 成功返回结果
type TaobaonextonelogisticssignupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"nextone_logistics_sign_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaonextonelogisticssignupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
