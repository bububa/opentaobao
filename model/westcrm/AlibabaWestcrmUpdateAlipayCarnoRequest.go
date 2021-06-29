package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新支付宝业务卡号 API请求
alibaba.westcrm.update.alipay.carno

更新支付宝业务卡号
*/
type AlibabaWestcrmUpdateAlipayCarnoRequest struct {
    model.Params
    // 商场id
    _mallId   int64
    // 用户id
    _id   int64
    // 2088102011918821
    _alipayCardNo   string
    // appkey
    _westcrmAppKey   string
}

// 初始化AlibabaWestcrmUpdateAlipayCarnoRequest对象
func NewAlibabaWestcrmUpdateAlipayCarnoRequest() *AlibabaWestcrmUpdateAlipayCarnoRequest{
    return &AlibabaWestcrmUpdateAlipayCarnoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetApiMethodName() string {
    return "alibaba.westcrm.update.alipay.carno"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MallId Setter
// 商场id
func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetMallId(_mallId int64) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetMallId() int64 {
    return r._mallId
}
// Id Setter
// 用户id
func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetId() int64 {
    return r._id
}
// AlipayCardNo Setter
// 2088102011918821
func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetAlipayCardNo(_alipayCardNo string) error {
    r._alipayCardNo = _alipayCardNo
    r.Set("alipay_card_no", _alipayCardNo)
    return nil
}

// AlipayCardNo Getter
func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetAlipayCardNo() string {
    return r._alipayCardNo
}
// WestcrmAppKey Setter
// appkey
func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetWestcrmAppKey(_westcrmAppKey string) error {
    r._westcrmAppKey = _westcrmAppKey
    r.Set("westcrm_app_key", _westcrmAppKey)
    return nil
}

// WestcrmAppKey Getter
func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetWestcrmAppKey() string {
    return r._westcrmAppKey
}