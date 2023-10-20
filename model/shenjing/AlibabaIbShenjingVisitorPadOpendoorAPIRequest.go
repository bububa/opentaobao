package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaibshenjingvisitorpadopendoorAPIRequest 访客发起开门 API请求
// alibaba.ib.shenjing.visitor.pad.opendoor
//
// 访客PAD端录入完人脸后，可以点击开门按钮开门。
type AlibabaibshenjingvisitorpadopendoorAPIRequest struct {
	model.Params
	// 访客标识
	_id string
	// padid
	_padId string
}

// NewAlibabaibshenjingvisitorpadopendoorRequest 初始化AlibabaibshenjingvisitorpadopendoorAPIRequest对象
func NewAlibabaibshenjingvisitorpadopendoorRequest() *AlibabaibshenjingvisitorpadopendoorAPIRequest {
	return &AlibabaibshenjingvisitorpadopendoorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaibshenjingvisitorpadopendoorAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.opendoor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaibshenjingvisitorpadopendoorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaibshenjingvisitorpadopendoorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 访客标识
func (r *AlibabaibshenjingvisitorpadopendoorAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaibshenjingvisitorpadopendoorAPIRequest) GetId() string {
	return r._id
}

// SetPadId is PadId Setter
// padid
func (r *AlibabaibshenjingvisitorpadopendoorAPIRequest) SetPadId(_padId string) error {
	r._padId = _padId
	r.Set("pad_id", _padId)
	return nil
}

// GetPadId PadId Getter
func (r AlibabaibshenjingvisitorpadopendoorAPIRequest) GetPadId() string {
	return r._padId
}
