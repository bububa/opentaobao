package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmUpdateAlipayCarnoAPIRequest 更新支付宝业务卡号 API请求
// alibaba.westcrm.update.alipay.carno
//
// 更新支付宝业务卡号
type AlibabaWestcrmUpdateAlipayCarnoAPIRequest struct {
	model.Params
	// 商场id
	_mallId int64
	// 用户id
	_id int64
	// 2088102011918821
	_alipayCardNo string
	// appkey
	_westcrmAppKey string
}

// NewAlibabaWestcrmUpdateAlipayCarnoRequest 初始化AlibabaWestcrmUpdateAlipayCarnoAPIRequest对象
func NewAlibabaWestcrmUpdateAlipayCarnoRequest() *AlibabaWestcrmUpdateAlipayCarnoAPIRequest {
	return &AlibabaWestcrmUpdateAlipayCarnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUpdateAlipayCarnoAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.update.alipay.carno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUpdateAlipayCarnoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MallId Setter
// 商场id
func (r *AlibabaWestcrmUpdateAlipayCarnoAPIRequest) SetMallId(_mallId int64) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// Get MallId Getter
func (r AlibabaWestcrmUpdateAlipayCarnoAPIRequest) GetMallId() int64 {
	return r._mallId
}

// Set is Id Setter
// 用户id
func (r *AlibabaWestcrmUpdateAlipayCarnoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaWestcrmUpdateAlipayCarnoAPIRequest) GetId() int64 {
	return r._id
}

// Set is AlipayCardNo Setter
// 2088102011918821
func (r *AlibabaWestcrmUpdateAlipayCarnoAPIRequest) SetAlipayCardNo(_alipayCardNo string) error {
	r._alipayCardNo = _alipayCardNo
	r.Set("alipay_card_no", _alipayCardNo)
	return nil
}

// Get AlipayCardNo Getter
func (r AlibabaWestcrmUpdateAlipayCarnoAPIRequest) GetAlipayCardNo() string {
	return r._alipayCardNo
}

// Set is WestcrmAppKey Setter
// appkey
func (r *AlibabaWestcrmUpdateAlipayCarnoAPIRequest) SetWestcrmAppKey(_westcrmAppKey string) error {
	r._westcrmAppKey = _westcrmAppKey
	r.Set("westcrm_app_key", _westcrmAppKey)
	return nil
}

// Get WestcrmAppKey Getter
func (r AlibabaWestcrmUpdateAlipayCarnoAPIRequest) GetWestcrmAppKey() string {
	return r._westcrmAppKey
}
