package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseordernewgetstatusAPIRequest 订单状态查询接口 API请求
// alibaba.ele.enterprise.ordernew.getstatus
//
// 订单状态查询接口
type AlibabaeleenterpriseordernewgetstatusAPIRequest struct {
	model.Params
	// 订单号
	_elemeOrderId string
}

// NewAlibabaeleenterpriseordernewgetstatusRequest 初始化AlibabaeleenterpriseordernewgetstatusAPIRequest对象
func NewAlibabaeleenterpriseordernewgetstatusRequest() *AlibabaeleenterpriseordernewgetstatusAPIRequest {
	return &AlibabaeleenterpriseordernewgetstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseordernewgetstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.getstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseordernewgetstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseordernewgetstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetElemeOrderId is ElemeOrderId Setter
// 订单号
func (r *AlibabaeleenterpriseordernewgetstatusAPIRequest) SetElemeOrderId(_elemeOrderId string) error {
	r._elemeOrderId = _elemeOrderId
	r.Set("eleme_order_id", _elemeOrderId)
	return nil
}

// GetElemeOrderId ElemeOrderId Getter
func (r AlibabaeleenterpriseordernewgetstatusAPIRequest) GetElemeOrderId() string {
	return r._elemeOrderId
}
