package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoredeliverconfiggetAPIResponse 查询门店发货配置内容 API返回值
// taobao.omniorder.store.deliverconfig.get
//
// 查询门店发货配置内容
type TaobaoomniorderstoredeliverconfiggetAPIResponse struct {
	model.CommonResponse
	TaobaoomniorderstoredeliverconfiggetAPIResponseModel
}

// TaobaoomniorderstoredeliverconfiggetAPIResponseModel is 查询门店发货配置内容 成功返回结果
type TaobaoomniorderstoredeliverconfiggetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_deliverconfig_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoomniorderstoredeliverconfiggetResult `json:"result,omitempty" xml:"result,omitempty"`
}
