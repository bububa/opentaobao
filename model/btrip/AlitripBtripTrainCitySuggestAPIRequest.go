package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripTrainCitySuggestAPIRequest
火车票城市搜索 API请求
alitrip.btrip.train.city.suggest

阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接 */
type AlitripBtripTrainCitySuggestAPIRequest struct {
	model.Params
	// 用户id
	_userId string
	// 搜索关键字
	_keyword string
	// 企业id
	_corpId string
}

// New
