package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripTrainCitySuggestAPIResponse 火车票城市搜索 API返回值
// alitrip.btrip.train.city.suggest
//
// 阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
type AlitripBtripTrainCitySuggestAPIResponse struct {
	model.CommonResponse
	AlitripBtripTrainCitySuggestAPIResponseModel
}

// AlitripBtripTrainCitySuggestAPIResponseModel is 火车票城市搜索 成功返回结果
type AlitripBtripTrainCitySuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_train_city_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
