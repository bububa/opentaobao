package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微博公众号价格变化通知 API请求
alibaba.pictures.dengta.wbaccount.price.change

微博公众号推广价格变更通知接口
*/
type AlibabaPicturesDengtaWbaccountPriceChangeRequest struct {
    model.Params
    // 账号id
    _accountId   string
    // 转发价格
    _transferPrice   string
    // 日期
    _changeTime   string
    // 原发价
    _originPrice   string
    // id
    _id   int64
    // 转发价格 折后价
    _transferPriceAli   string
    // 原发价 折后价
    _originPriceAli   string
}

// 初始化AlibabaPicturesDengtaWbaccountPriceChangeRequest对象
func NewAlibabaPicturesDengtaWbaccountPriceChangeRequest() *AlibabaPicturesDengtaWbaccountPriceChangeRequest{
    return &AlibabaPicturesDengtaWbaccountPriceChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.wbaccount.price.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 账号id
func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetAccountId() string {
    return r._accountId
}
// TransferPrice Setter
// 转发价格
func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetTransferPrice(_transferPrice string) error {
    r._transferPrice = _transferPrice
    r.Set("transfer_price", _transferPrice)
    return nil
}

// TransferPrice Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetTransferPrice() string {
    return r._transferPrice
}
// ChangeTime Setter
// 日期
func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetChangeTime(_changeTime string) error {
    r._changeTime = _changeTime
    r.Set("change_time", _changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetChangeTime() string {
    return r._changeTime
}
// OriginPrice Setter
// 原发价
func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetOriginPrice(_originPrice string) error {
    r._originPrice = _originPrice
    r.Set("origin_price", _originPrice)
    return nil
}

// OriginPrice Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetOriginPrice() string {
    return r._originPrice
}
// Id Setter
// id
func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetId() int64 {
    return r._id
}
// TransferPriceAli Setter
// 转发价格 折后价
func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetTransferPriceAli(_transferPriceAli string) error {
    r._transferPriceAli = _transferPriceAli
    r.Set("transfer_price_ali", _transferPriceAli)
    return nil
}

// TransferPriceAli Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetTransferPriceAli() string {
    return r._transferPriceAli
}
// OriginPriceAli Setter
// 原发价 折后价
func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetOriginPriceAli(_originPriceAli string) error {
    r._originPriceAli = _originPriceAli
    r.Set("origin_price_ali", _originPriceAli)
    return nil
}

// OriginPriceAli Getter
func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetOriginPriceAli() string {
    return r._originPriceAli
}
