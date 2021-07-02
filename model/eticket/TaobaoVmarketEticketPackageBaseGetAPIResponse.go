package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketPackageBaseGetAPIResponse 获取包基本信息 API返回值
// taobao.vmarket.eticket.package.base.get
//
// 获取包基本信息
type TaobaoVmarketEticketPackageBaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketPackageBaseGetAPIResponseModel
}

// TaobaoVmarketEticketPackageBaseGetAPIResponseModel is 获取包基本信息 成功返回结果
type TaobaoVmarketEticketPackageBaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_package_base_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *PackageResult `json:"result,omitempty" xml:"result,omitempty"`
}
