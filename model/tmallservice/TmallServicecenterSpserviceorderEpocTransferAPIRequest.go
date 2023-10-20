package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterspserviceorderepoctransferAPIRequest 电子保单受保人转移 API请求
// tmall.servicecenter.spserviceorder.epoc.transfer
//
// 电子保单受保人转移
type TmallservicecenterspserviceorderepoctransferAPIRequest struct {
	model.Params
	// 服务交易订单号
	_bizOrderId int64
}

// NewTmallservicecenterspserviceorderepoctransferRequest 初始化TmallservicecenterspserviceorderepoctransferAPIRequest对象
func NewTmallservicecenterspserviceorderepoctransferRequest() *TmallservicecenterspserviceorderepoctransferAPIRequest {
	return &TmallservicecenterspserviceorderepoctransferAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterspserviceorderepoctransferAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.epoc.transfer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterspserviceorderepoctransferAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterspserviceorderepoctransferAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 服务交易订单号
func (r *TmallservicecenterspserviceorderepoctransferAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallservicecenterspserviceorderepoctransferAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
