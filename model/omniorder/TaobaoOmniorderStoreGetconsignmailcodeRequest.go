package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店物流菜鸟裹裹取号 APIRequest
taobao.omniorder.store.getconsignmailcode

用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
*/
type TaobaoOmniorderStoreGetconsignmailcodeRequest struct {
    model.Params

    // 门店ID
    storeId   int64 

    // 发件人联系电话，如空则表示使用门店信息中的电话号码
    senderContact   string 

    // 淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)      、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)      、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里      巴巴(1688)、其他(OTHERS)
    channel   string 

    // 订单信息，目前一次请求只支持一个主订单
    trades   []TradeOrderInfoDto 

    // 收件人信息
    receiver   *ReceiverDto 

    // 扩展信息
    sdtExtendInfoDTO   *SdtExtendInfoDto 

}

func NewTaobaoOmniorderStoreGetconsignmailcodeRequest() *TaobaoOmniorderStoreGetconsignmailcodeRequest{
    return &TaobaoOmniorderStoreGetconsignmailcodeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.getconsignmailcode"
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetSenderContact(senderContact string) error {
    r.senderContact = senderContact
    r.Set("sender_contact", senderContact)
    return nil
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetSenderContact() string {
    return r.senderContact
}

func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetChannel() string {
    return r.channel
}

func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetTrades(trades []TradeOrderInfoDto) error {
    r.trades = trades
    r.Set("trades", trades)
    return nil
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetTrades() []TradeOrderInfoDto {
    return r.trades
}

func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetReceiver(receiver *ReceiverDto) error {
    r.receiver = receiver
    r.Set("receiver", receiver)
    return nil
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetReceiver() *ReceiverDto {
    return r.receiver
}

func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetSdtExtendInfoDTO(sdtExtendInfoDTO *SdtExtendInfoDto) error {
    r.sdtExtendInfoDTO = sdtExtendInfoDTO
    r.Set("sdt_extend_info_d_t_o", sdtExtendInfoDTO)
    return nil
}

func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetSdtExtendInfoDTO() *SdtExtendInfoDto {
    return r.sdtExtendInfoDTO
}

