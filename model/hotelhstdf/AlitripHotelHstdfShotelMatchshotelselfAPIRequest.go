package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfshotelmatchshotelselfAPIRequest 自主匹配标准酒店以及卖家酒店 API请求
// alitrip.hotel.hstdf.shotel.matchshotelself
//
// 商家通过指定的标准酒店id和卖家酒店id进行匹配
type AlitriphotelhstdfshotelmatchshotelselfAPIRequest struct {
	model.Params
	// HotelMatchParam
	_param0 *HotelMatchParam
}

// NewAlitriphotelhstdfshotelmatchshotelselfRequest 初始化AlitriphotelhstdfshotelmatchshotelselfAPIRequest对象
func NewAlitriphotelhstdfshotelmatchshotelselfRequest() *AlitriphotelhstdfshotelmatchshotelselfAPIRequest {
	return &AlitriphotelhstdfshotelmatchshotelselfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfshotelmatchshotelselfAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.matchshotelself"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfshotelmatchshotelselfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfshotelmatchshotelselfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// HotelMatchParam
func (r *AlitriphotelhstdfshotelmatchshotelselfAPIRequest) SetParam0(_param0 *HotelMatchParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitriphotelhstdfshotelmatchshotelselfAPIRequest) GetParam0() *HotelMatchParam {
	return r._param0
}
