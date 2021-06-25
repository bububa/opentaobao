package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买家申请的退款列表 APIRequest
taobao.refunds.apply.get

查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
*/
type TaobaoRefundsApplyGetRequest struct {
    model.Params

    // 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee
    fields   []String 

    // 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。<br/>WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) <br/>WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) <br/>WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) <br/>SELLER_REFUSE_BUYER(卖家拒绝退款) <br/>CLOSED(退款关闭) <br/>SUCCESS(退款成功)
    status   string 

    // 卖家昵称
    sellerNick   string 

    // 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。<br/>fixed(一口价) <br/>auction(拍卖) <br/>guarantee_trade(一口价、拍卖) <br/>independent_simple_trade(旺店入门版交易) <br/>independent_shop_trade(旺店标准版交易) <br/>auto_delivery(自动发货) <br/>ec(直冲) <br/>cod(货到付款) <br/>fenxiao(分销) <br/>game_equipment(游戏装备) <br/>shopex_trade(ShopEX交易) <br/>netcn_trade(万网交易) <br/>external_trade(统一外部交易)
    type   string 

    // 页码。传入值为 1 代表第一页，传入值为 2 代表第二页，依此类推。默认返回的数据是从第一页开始
    pageNo   int64 

    // 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
    pageSize   int64 

}

func NewTaobaoRefundsApplyGetRequest() *TaobaoRefundsApplyGetRequest{
    return &TaobaoRefundsApplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRefundsApplyGetRequest) GetApiMethodName() string {
    return "taobao.refunds.apply.get"
}

func (r TaobaoRefundsApplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRefundsApplyGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoRefundsApplyGetRequest) GetFields() []String {
    return r.fields
}

func (r *TaobaoRefundsApplyGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoRefundsApplyGetRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoRefundsApplyGetRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r TaobaoRefundsApplyGetRequest) GetSellerNick() string {
    return r.sellerNick
}

func (r *TaobaoRefundsApplyGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoRefundsApplyGetRequest) GetType() string {
    return r.type
}

func (r *TaobaoRefundsApplyGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoRefundsApplyGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoRefundsApplyGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoRefundsApplyGetRequest) GetPageSize() int64 {
    return r.pageSize
}

