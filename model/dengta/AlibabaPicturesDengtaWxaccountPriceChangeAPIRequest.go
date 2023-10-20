package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapicturesdengtawxaccountpricechangeAPIRequest 微信公众号价格变化通知 API请求
// alibaba.pictures.dengta.wxaccount.price.change
//
// 微信公众号推广价格变更通知接口
type AlibabapicturesdengtawxaccountpricechangeAPIRequest struct {
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

// NewAlibabapicturesdengtawxaccountpricechangeRequest 初始化AlibabapicturesdengtawxaccountpricechangeAPIRequest对象
func NewAlibabapicturesdengtawxaccountpricechangeRequest() *AlibabapicturesdengtawxaccountpricechangeAPIRequest {
	return &AlibabapicturesdengtawxaccountpricechangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.wxaccount.price.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// id
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetChangeTime is ChangeTime Setter
// 变更时间
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetSingle is Single Setter
// 单图文
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetSingle(_single string) error {
	r._single = _single
	r.Set("single", _single)
	return nil
}

// GetSingle Single Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetSingle() string {
	return r._single
}

// SetOther is Other Setter
// 多图文第三条及以后
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetOther(_other string) error {
	r._other = _other
	r.Set("other", _other)
	return nil
}

// GetOther Other Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetOther() string {
	return r._other
}

// SetSecond is Second Setter
// 多图文第二条
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetSecond(_second string) error {
	r._second = _second
	r.Set("second", _second)
	return nil
}

// GetSecond Second Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetSecond() string {
	return r._second
}

// SetFirst is First Setter
// 多图文第一条
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetFirst(_first string) error {
	r._first = _first
	r.Set("first", _first)
	return nil
}

// GetFirst First Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetFirst() string {
	return r._first
}

// SetFirstAli is FirstAli Setter
// 多图文第一条 折后价
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetFirstAli(_firstAli string) error {
	r._firstAli = _firstAli
	r.Set("first_ali", _firstAli)
	return nil
}

// GetFirstAli FirstAli Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetFirstAli() string {
	return r._firstAli
}

// SetSecondAli is SecondAli Setter
// 多图文第二条 折后价
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetSecondAli(_secondAli string) error {
	r._secondAli = _secondAli
	r.Set("second_ali", _secondAli)
	return nil
}

// GetSecondAli SecondAli Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetSecondAli() string {
	return r._secondAli
}

// SetSingleAli is SingleAli Setter
// 单图文 折后价
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetSingleAli(_singleAli string) error {
	r._singleAli = _singleAli
	r.Set("single_ali", _singleAli)
	return nil
}

// GetSingleAli SingleAli Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetSingleAli() string {
	return r._singleAli
}

// SetOtherAli is OtherAli Setter
// 多图文第三条及以后 折后价
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetOtherAli(_otherAli string) error {
	r._otherAli = _otherAli
	r.Set("other_ali", _otherAli)
	return nil
}

// GetOtherAli OtherAli Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetOtherAli() string {
	return r._otherAli
}

// SetId is Id Setter
// 账号id
func (r *AlibabapicturesdengtawxaccountpricechangeAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabapicturesdengtawxaccountpricechangeAPIRequest) GetId() int64 {
	return r._id
}
