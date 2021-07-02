package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest 查询当天可添加的余量 API请求
// taobao.baichuan.item.subscribe.daily.left.query
//
// 查询当天可添加的余量
type TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest struct {
	model.Params
}

// NewTaobaoBaichuanItemSubscribeDailyLeftQueryRequest 初始化TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest对象
func NewTaobaoBaichuanItemSubscribeDailyLeftQueryRequest() *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest {
	return &TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe.daily.left.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
