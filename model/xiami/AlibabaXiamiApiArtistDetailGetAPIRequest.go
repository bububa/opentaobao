package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiArtistDetailGetAPIRequest 艺人详情 API请求
// alibaba.xiami.api.artist.detail.get
//
// 艺人详情
type AlibabaXiamiApiArtistDetailGetAPIRequest struct {
	model.Params
	// 艺人id
	_id int64
	// 是否显示description, show为显示, 其他为不显示
	_description string
}

// NewAlibabaXiamiApiArtistDetailGetRequest 初始化AlibabaXiamiApiArtistDetailGetAPIRequest对象
func NewAlibabaXiamiApiArtistDetailGetRequest() *AlibabaXiamiApiArtistDetailGetAPIRequest {
	return &AlibabaXiamiApiArtistDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.artist.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetId is Id Setter
// 艺人id
func (r *AlibabaXiamiApiArtistDetailGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaXiamiApiArtistDetailGetAPIRequest) GetId() int64 {
	return r._id
}

// SetDescription is Description Setter
// 是否显示description, show为显示, 其他为不显示
func (r *AlibabaXiamiApiArtistDetailGetAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r AlibabaXiamiApiArtistDetailGetAPIRequest) GetDescription() string {
	return r._description
}
