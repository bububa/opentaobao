package dengta

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest 微信公众号价格变化通知 API请求
// alibaba.pictures.dengta.wxaccount.price.change
//
// 微信公众号推广价格变更通知接口
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
	// 多图文第一条 折后价
	_firstAli string
	// 多图文第二条 折后价
	_secondAli string
	// 单图文 折后价
	_singleAli string
	// 多图文第三条及以后 折后价
	_otherAli string
	// 账号id
	_id int64
}

// NewAlibabaPicturesDengtaWxaccountPriceChangeRequest 初始化AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest对象
func NewAlibabaPicturesDengtaWxaccountPriceChangeRequest() *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest {
	return &AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) Reset() {
	r._accountId = ""
	r._changeTime = ""
	r._single = ""
	r._other = ""
	r._second = ""
	r._first = ""
	r._firstAli = ""
	r._secondAli = ""
	r._singleAli = ""
	r._otherAli = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.wxaccount.price.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// id
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetChangeTime is ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetSingle is Single Setter
// 单图文
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSingle(_single string) error {
	r._single = _single
	r.Set("single", _single)
	return nil
}

// GetSingle Single Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSingle() string {
	return r._single
}

// SetOther is Other Setter
// 多图文第三条及以后
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetOther(_other string) error {
	r._other = _other
	r.Set("other", _other)
	return nil
}

// GetOther Other Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetOther() string {
	return r._other
}

// SetSecond is Second Setter
// 多图文第二条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSecond(_second string) error {
	r._second = _second
	r.Set("second", _second)
	return nil
}

// GetSecond Second Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSecond() string {
	return r._second
}

// SetFirst is First Setter
// 多图文第一条
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetFirst(_first string) error {
	r._first = _first
	r.Set("first", _first)
	return nil
}

// GetFirst First Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetFirst() string {
	return r._first
}

// SetFirstAli is FirstAli Setter
// 多图文第一条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetFirstAli(_firstAli string) error {
	r._firstAli = _firstAli
	r.Set("first_ali", _firstAli)
	return nil
}

// GetFirstAli FirstAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetFirstAli() string {
	return r._firstAli
}

// SetSecondAli is SecondAli Setter
// 多图文第二条 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSecondAli(_secondAli string) error {
	r._secondAli = _secondAli
	r.Set("second_ali", _secondAli)
	return nil
}

// GetSecondAli SecondAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSecondAli() string {
	return r._secondAli
}

// SetSingleAli is SingleAli Setter
// 单图文 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetSingleAli(_singleAli string) error {
	r._singleAli = _singleAli
	r.Set("single_ali", _singleAli)
	return nil
}

// GetSingleAli SingleAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetSingleAli() string {
	return r._singleAli
}

// SetOtherAli is OtherAli Setter
// 多图文第三条及以后 折后价
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetOtherAli(_otherAli string) error {
	r._otherAli = _otherAli
	r.Set("other_ali", _otherAli)
	return nil
}

// GetOtherAli OtherAli Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetOtherAli() string {
	return r._otherAli
}

// SetId is Id Setter
// 账号id
func (r *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) GetId() int64 {
	return r._id
}

var poolAlibabaPicturesDengtaWxaccountPriceChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPicturesDengtaWxaccountPriceChangeRequest()
	},
}

// GetAlibabaPicturesDengtaWxaccountPriceChangeRequest 从 sync.Pool 获取 AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest
func GetAlibabaPicturesDengtaWxaccountPriceChangeAPIRequest() *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest {
	return poolAlibabaPicturesDengtaWxaccountPriceChangeAPIRequest.Get().(*AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest)
}

// ReleaseAlibabaPicturesDengtaWxaccountPriceChangeAPIRequest 将 AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaPicturesDengtaWxaccountPriceChangeAPIRequest(v *AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest) {
	v.Reset()
	poolAlibabaPicturesDengtaWxaccountPriceChangeAPIRequest.Put(v)
}
