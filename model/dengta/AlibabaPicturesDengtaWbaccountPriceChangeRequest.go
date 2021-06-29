package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微博公众号价格变化通知 APIRequest
alibaba.pictures.dengta.wbaccount.price.change

微博公众号推广价格变更通知接口
*/
type AlibabaPicturesDengtaWbaccountPriceChangeRequest struct {
    model.Params

    // 账号id
    accountId   string 

    // 转发价格
    transferPrice   string 

    // 日期
    changeTime   string 

    // 原发价
    originPrice   string 

    // id
    id   int64 

    // 转发价格 折后价
    transferPriceAli   string 

    // 原发价 折后价
    originPriceAli   string 

}

func NewAlibabaPicturesDengtaWbaccountPriceChangeRequest() *AlibabaPicturesDengtaWbaccountPriceChangeRequest{
    return &AlibabaPicturesDengtaWbaccountPriceChangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.wbaccount.price.change"
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetTransferPrice(transferPrice string) error {
    r.transferPrice = transferPrice
    r.Set("transfer_price", transferPrice)
    return nil
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetTransferPrice() string {
    return r.transferPrice
}

func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetChangeTime(changeTime string) error {
    r.changeTime = changeTime
    r.Set("change_time", changeTime)
    return nil
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetChangeTime() string {
    return r.changeTime
}

func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetOriginPrice(originPrice string) error {
    r.originPrice = originPrice
    r.Set("origin_price", originPrice)
    return nil
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetOriginPrice() string {
    return r.originPrice
}

func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetTransferPriceAli(transferPriceAli string) error {
    r.transferPriceAli = transferPriceAli
    r.Set("transfer_price_ali", transferPriceAli)
    return nil
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetTransferPriceAli() string {
    return r.transferPriceAli
}

func (r *AlibabaPicturesDengtaWbaccountPriceChangeRequest) SetOriginPriceAli(originPriceAli string) error {
    r.originPriceAli = originPriceAli
    r.Set("origin_price_ali", originPriceAli)
    return nil
}

func (r AlibabaPicturesDengtaWbaccountPriceChangeRequest) GetOriginPriceAli() string {
    return r.originPriceAli
}

