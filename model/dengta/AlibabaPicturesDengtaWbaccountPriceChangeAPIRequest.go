package dengta

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest 微博公众号价格变化通知 API请求
// alibaba.pictures.dengta.wbaccount.price.change
//
// 微博公众号推广价格变更通知接口
type AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest struct {
	model.Params
	// 账号id
	_accountId string
	// 转发价格
	_transferPrice string
	// 日期
	_changeTime string
	// 转发价格 折后价
	_transferPriceAli string
	// 原发价
	_originPrice string
	// 原发价 折后价
	_originPriceAli string
	// id
	_id int64
}

// NewAlibabaPicturesDengtaWbaccountPriceChangeRequest 初始化AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest对象
func NewAlibabaPicturesDengtaWbaccountPriceChangeRequest() *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest {
	return &AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) Reset() {
	r._accountId = ""
	r._transferPrice = ""
	r._changeTime = ""
	r._transferPriceAli = ""
	r._originPrice = ""
	r._originPriceAli = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.wbaccount.price.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 账号id
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetTransferPrice is TransferPrice Setter
// 转发价格
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) SetTransferPrice(_transferPrice string) error {
	r._transferPrice = _transferPrice
	r.Set("transfer_price", _transferPrice)
	return nil
}

// GetTransferPrice TransferPrice Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetTransferPrice() string {
	return r._transferPrice
}

// SetChangeTime is ChangeTime Setter
// 日期
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetTransferPriceAli is TransferPriceAli Setter
// 转发价格 折后价
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) SetTransferPriceAli(_transferPriceAli string) error {
	r._transferPriceAli = _transferPriceAli
	r.Set("transfer_price_ali", _transferPriceAli)
	return nil
}

// GetTransferPriceAli TransferPriceAli Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetTransferPriceAli() string {
	return r._transferPriceAli
}

// SetOriginPrice is OriginPrice Setter
// 原发价
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) SetOriginPrice(_originPrice string) error {
	r._originPrice = _originPrice
	r.Set("origin_price", _originPrice)
	return nil
}

// GetOriginPrice OriginPrice Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetOriginPrice() string {
	return r._originPrice
}

// SetOriginPriceAli is OriginPriceAli Setter
// 原发价 折后价
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) SetOriginPriceAli(_originPriceAli string) error {
	r._originPriceAli = _originPriceAli
	r.Set("origin_price_ali", _originPriceAli)
	return nil
}

// GetOriginPriceAli OriginPriceAli Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetOriginPriceAli() string {
	return r._originPriceAli
}

// SetId is Id Setter
// id
func (r *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) GetId() int64 {
	return r._id
}

var poolAlibabaPicturesDengtaWbaccountPriceChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPicturesDengtaWbaccountPriceChangeRequest()
	},
}

// GetAlibabaPicturesDengtaWbaccountPriceChangeRequest 从 sync.Pool 获取 AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest
func GetAlibabaPicturesDengtaWbaccountPriceChangeAPIRequest() *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest {
	return poolAlibabaPicturesDengtaWbaccountPriceChangeAPIRequest.Get().(*AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest)
}

// ReleaseAlibabaPicturesDengtaWbaccountPriceChangeAPIRequest 将 AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaPicturesDengtaWbaccountPriceChangeAPIRequest(v *AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest) {
	v.Reset()
	poolAlibabaPicturesDengtaWbaccountPriceChangeAPIRequest.Put(v)
}
