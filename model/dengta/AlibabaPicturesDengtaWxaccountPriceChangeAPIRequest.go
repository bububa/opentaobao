package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest
微信公众号价格变化通知 API请求
alibaba.pictures.dengta.wxaccount.price.change

微信公众号推广价格变更通知接口 */
type AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest struct {
	model.Params
	// id
	_accountId string
	// 变更时间
	_changeTime string
	// 单图文
	_single string
	// 多图文第三条及以后
	_other string
	// 多图文第二条
	_second string
	// 多图文第一条
	_first string
	// 账号id
	_id int64
	// 多图文第一条 折后价
	_firstAli string
	// 多图文第二条 折后价
	_secondAli string
	// 单图文 折后价
	_singleAli string
	// 多图文第三条及以后 折后价
	_otherAli string
}

// NewAlibabaPicturesDengtaWxaccountPriceChangeRequest 初始化AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest对象
func NewAlibabaPicturesDengtaWxaccountPriceChangeRequest() *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest {
	return &AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.wxaccount.price.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AccountId Setter
// id
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// Get AccountId Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetAccountId() string {
	return r._accountId
}

// Set is ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// Get ChangeTime Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// Set is Single Setter
// 单图文
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSingle(_single string) error {
	r._single = _single
	r.Set("single", _single)
	return nil
}

// Get Single Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSingle() string {
	return r._single
}

// Set is Other Setter
// 多图文第三条及以后
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetOther(_other string) error {
	r._other = _other
	r.Set("other", _other)
	return nil
}

// Get Other Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetOther() string {
	return r._other
}

// Set is Second Setter
// 多图文第二条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSecond(_second string) error {
	r._second = _second
	r.Set("second", _second)
	return nil
}

// Get Second Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSecond() string {
	return r._second
}

// Set is First Setter
// 多图文第一条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetFirst(_first string) error {
	r._first = _first
	r.Set("first", _first)
	return nil
}

// Get First Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetFirst() string {
	return r._first
}

// Set is Id Setter
// 账号id
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetId() int64 {
	return r._id
}

// Set is FirstAli Setter
// 多图文第一条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetFirstAli(_firstAli string) error {
	r._firstAli = _firstAli
	r.Set("first_ali", _firstAli)
	return nil
}

// Get FirstAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetFirstAli() string {
	return r._firstAli
}

// Set is SecondAli Setter
// 多图文第二条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSecondAli(_secondAli string) error {
	r._secondAli = _secondAli
	r.Set("second_ali", _secondAli)
	return nil
}

// Get SecondAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSecondAli() string {
	return r._secondAli
}

// Set is SingleAli Setter
// 单图文 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSingleAli(_singleAli string) error {
	r._singleAli = _singleAli
	r.Set("single_ali", _singleAli)
	return nil
}

// Get SingleAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSingleAli() string {
	return r._singleAli
}

// Set is OtherAli Setter
// 多图文第三条及以后 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetOtherAli(_otherAli string) error {
	r._otherAli = _otherAli
	r.Set("other_ali", _otherAli)
	return nil
}

// Get OtherAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetOtherAli() string {
	return r._otherAli
}
