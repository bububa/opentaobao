package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomevirtualshopsyncAPIResponse 二手房虚拟店铺数据同步 API返回值
// alibaba.alihouse.existinghome.virtualshop.sync
//
// 二手房虚拟店铺数据同步
type AlibabaalihouseexistinghomevirtualshopsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomevirtualshopsyncAPIResponseModel
}

// AlibabaalihouseexistinghomevirtualshopsyncAPIResponseModel is 二手房虚拟店铺数据同步 成功返回结果
type AlibabaalihouseexistinghomevirtualshopsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_virtualshop_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomevirtualshopsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
