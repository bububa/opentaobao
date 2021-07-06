package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest 阿里实人认证卡片信息回传 API请求
// alibaba.alitj.order.realnamecard.info.submit
//
// 阿里实人认证卡片信息回传。ISP相关商家在线对接阿里通信的实人认证功能，在线提交订单对应运营商的合约订购相关信息，以便完成在线使用实人认证功能。
type AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest struct {
	model.Params
	// sim卡iccid（一般为18位到20位）
	_iccid string
	// 淘宝订单号
	_orderNo int64
}

// NewAlibabaAlitjOrderRealnamecardInfoSubmitRequest 初始化AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest对象
func NewAlibabaAlitjOrderRealnamecardInfoSubmitRequest() *AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest {
	return &AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alitj.order.realnamecard.info.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIccid is Iccid Setter
// sim卡iccid（一般为18位到20位）
func (r *AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest) GetIccid() string {
	return r._iccid
}

// SetOrderNo is OrderNo Setter
// 淘宝订单号
func (r *AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest) SetOrderNo(_orderNo int64) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest) GetOrderNo() int64 {
	return r._orderNo
}
