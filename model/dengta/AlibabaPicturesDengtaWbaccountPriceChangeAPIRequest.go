package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapicturesdengtawbaccountpricechangeAPIRequest 微博公众号价格变化通知 API请求
// alibaba.pictures.dengta.wbaccount.price.change
//
// 微博公众号推广价格变更通知接口
type AlibabapicturesdengtawbaccountpricechangeAPIRequest struct {
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

// NewAlibabapicturesdengtawbaccountpricechangeRequest 初始化AlibabapicturesdengtawbaccountpricechangeAPIRequest对象
func NewAlibabapicturesdengtawbaccountpricechangeRequest() *AlibabapicturesdengtawbaccountpricechangeAPIRequest {
	return &AlibabapicturesdengtawbaccountpricechangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.wbaccount.price.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 账号id
func (r *AlibabapicturesdengtawbaccountpricechangeAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetTransferPrice is TransferPrice Setter
// 转发价格
func (r *AlibabapicturesdengtawbaccountpricechangeAPIRequest) SetTransferPrice(_transferPrice string) error {
	r._transferPrice = _transferPrice
	r.Set("transfer_price", _transferPrice)
	return nil
}

// GetTransferPrice TransferPrice Getter
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetTransferPrice() string {
	return r._transferPrice
}

// SetChangeTime is ChangeTime Setter
// 日期
func (r *AlibabapicturesdengtawbaccountpricechangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetTransferPriceAli is TransferPriceAli Setter
// 转发价格 折后价
func (r *AlibabapicturesdengtawbaccountpricechangeAPIRequest) SetTransferPriceAli(_transferPriceAli string) error {
	r._transferPriceAli = _transferPriceAli
	r.Set("transfer_price_ali", _transferPriceAli)
	return nil
}

// GetTransferPriceAli TransferPriceAli Getter
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetTransferPriceAli() string {
	return r._transferPriceAli
}

// SetOriginPrice is OriginPrice Setter
// 原发价
func (r *AlibabapicturesdengtawbaccountpricechangeAPIRequest) SetOriginPrice(_originPrice string) error {
	r._originPrice = _originPrice
	r.Set("origin_price", _originPrice)
	return nil
}

// GetOriginPrice OriginPrice Getter
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetOriginPrice() string {
	return r._originPrice
}

// SetOriginPriceAli is OriginPriceAli Setter
// 原发价 折后价
func (r *AlibabapicturesdengtawbaccountpricechangeAPIRequest) SetOriginPriceAli(_originPriceAli string) error {
	r._originPriceAli = _originPriceAli
	r.Set("origin_price_ali", _originPriceAli)
	return nil
}

// GetOriginPriceAli OriginPriceAli Getter
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetOriginPriceAli() string {
	return r._originPriceAli
}

// SetId is Id Setter
// id
func (r *AlibabapicturesdengtawbaccountpricechangeAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabapicturesdengtawbaccountpricechangeAPIRequest) GetId() int64 {
	return r._id
}
