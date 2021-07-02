package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadOpendoorAPIRequest 访客发起开门 API请求
// alibaba.ib.shenjing.visitor.pad.opendoor
//
// 访客PAD端录入完人脸后，可以点击开门按钮开门。
type AlibabaIbShenjingVisitorPadOpendoorAPIRequest struct {
	model.Params
	// 访客标识
	_id string
	// padid
	_padId string
}

// NewAlibabaIbShenjingVisitorPadOpendoorRequest 初始化AlibabaIbShenjingVisitorPadOpendoorAPIRequest对象
func NewAlibabaIbShenjingVisitorPadOpendoorRequest() *AlibabaIbShenjingVisitorPadOpendoorAPIRequest {
	return &AlibabaIbShenjingVisitorPadOpendoorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadOpendoorAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.opendoor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadOpendoorAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetId is Id Setter
// 访客标识
func (r *AlibabaIbShenjingVisitorPadOpendoorAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaIbShenjingVisitorPadOpendoorAPIRequest) GetId() string {
	return r._id
}

// SetPadId is PadId Setter
// padid
func (r *AlibabaIbShenjingVisitorPadOpendoorAPIRequest) SetPadId(_padId string) error {
	r._padId = _padId
	r.Set("pad_id", _padId)
	return nil
}

// GetPadId PadId Getter
func (r AlibabaIbShenjingVisitorPadOpendoorAPIRequest) GetPadId() string {
	return r._padId
}
