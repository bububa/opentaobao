package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountUpdateAPIResponse Open Account数据更新 API返回值
// taobao.open.account.update
//
// Open Account数据更新
type TaobaoOpenAccountUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountUpdateAPIResponseModel
}

// TaobaoOpenAccountUpdateAPIResponseModel is Open Account数据更新 成功返回结果
type TaobaoOpenAccountUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// update是否成功
	Datas []OpenaccountVoid `json:"datas,omitempty" xml:"datas>openaccount_void,omitempty"`
}
