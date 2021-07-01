package store

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoregroupCreateAPIResponse
商户门店库创建接口 API返回值
taobao.place.storegroup.create

用于商家创建线下门店库 */
type TaobaoPlaceStoregroupCreateAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoregroupCreateAPIResponseModel
}

// TaobaoPlaceStoregroupCreateAPIResponseModel is 商户门店库创建接口 成功返回结果
type TaobaoPlaceStoregroupCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storegroup_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
