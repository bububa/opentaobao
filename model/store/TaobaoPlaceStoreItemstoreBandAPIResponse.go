package store

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreItemstoreBandAPIResponse 门店商品关联绑定接口 API返回值
// taobao.place.store.itemstore.band
//
// 商品和多个门店关系绑定接口
type TaobaoPlaceStoreItemstoreBandAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreItemstoreBandAPIResponseModel
}

// TaobaoPlaceStoreItemstoreBandAPIResponseModel is 门店商品关联绑定接口 成功返回结果
type TaobaoPlaceStoreItemstoreBandAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_itemstore_band_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
