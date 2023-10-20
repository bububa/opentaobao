package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptraveltradeserviceinfowriteAPIRequest 订单服务信息写入接口 API请求
// alitrip.travel.trade.serviceinfo.write
//
// 订单服务信息写入接口
type AlitriptraveltradeserviceinfowriteAPIRequest struct {
	model.Params
	// 根据模版要求传递对应的订单服务信息，当涉及多值时，用英文分号隔开";"，目前api暂时不支持文件上传，只支持链接；链接必须外网能访问
	_tipValue string
	// 对应的订单信息
	_subTcOrderId int64
}

// NewAlitriptraveltradeserviceinfowriteRequest 初始化AlitriptraveltradeserviceinfowriteAPIRequest对象
func NewAlitriptraveltradeserviceinfowriteRequest() *AlitriptraveltradeserviceinfowriteAPIRequest {
	return &AlitriptraveltradeserviceinfowriteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptraveltradeserviceinfowriteAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.serviceinfo.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptraveltradeserviceinfowriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptraveltradeserviceinfowriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTipValue is TipValue Setter
// 根据模版要求传递对应的订单服务信息，当涉及多值时，用英文分号隔开&#34;;&#34;，目前api暂时不支持文件上传，只支持链接；链接必须外网能访问
func (r *AlitriptraveltradeserviceinfowriteAPIRequest) SetTipValue(_tipValue string) error {
	r._tipValue = _tipValue
	r.Set("tip_value", _tipValue)
	return nil
}

// GetTipValue TipValue Getter
func (r AlitriptraveltradeserviceinfowriteAPIRequest) GetTipValue() string {
	return r._tipValue
}

// SetSubTcOrderId is SubTcOrderId Setter
// 对应的订单信息
func (r *AlitriptraveltradeserviceinfowriteAPIRequest) SetSubTcOrderId(_subTcOrderId int64) error {
	r._subTcOrderId = _subTcOrderId
	r.Set("sub_tc_order_id", _subTcOrderId)
	return nil
}

// GetSubTcOrderId SubTcOrderId Getter
func (r AlitriptraveltradeserviceinfowriteAPIRequest) GetSubTcOrderId() int64 {
	return r._subTcOrderId
}
