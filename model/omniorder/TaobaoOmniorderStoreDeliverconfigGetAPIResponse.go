package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreDeliverconfigGetAPIResponse
查询门店发货配置内容 API返回值
taobao.omniorder.store.deliverconfig.get

查询门店发货配置内容 */
type TaobaoOmniorderStoreDeliverconfigGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel
}

// TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel is 查询门店发货配置内容 成功返回结果
type TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_deliverconfig_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreDeliverconfigGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
