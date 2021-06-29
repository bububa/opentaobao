package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店物流菜鸟裹裹取号 API请求
taobao.omniorder.store.getconsignmailcode

用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
*/
type TaobaoOmniorderStoreGetconsignmailcodeRequest struct {
    model.Params
    // 门店ID
    _storeId   int64
    // 发件人联系电话，如空则表示使用门店信息中的电话号码
    _senderContact   string
    // 淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)      、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)      、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里      巴巴(1688)、其他(OTHERS)
    _channel   string
    // 订单信息，目前一次请求只支持一个主订单
    _trades   []TradeOrderInfoDto
    // 收件人信息
    _receiver   *ReceiverDto
    // 扩展信息
    _sdtExtendInfoDTO   *SdtExtendInfoDto
}

// 初始化TaobaoOmniorderStoreGetconsignmailcodeRequest对象
func NewTaobaoOmniorderStoreGetconsignmailcodeRequest() *TaobaoOmniorderStoreGetconsignmailcodeRequest{
    return &TaobaoOmniorderStoreGetconsignmailcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.getconsignmailcode"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetStoreId() int64 {
    return r._storeId
}
// SenderContact Setter
// 发件人联系电话，如空则表示使用门店信息中的电话号码
func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetSenderContact(_senderContact string) error {
    r._senderContact = _senderContact
    r.Set("sender_contact", _senderContact)
    return nil
}

// SenderContact Getter
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetSenderContact() string {
    return r._senderContact
}
// Channel Setter
// 淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)      、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)      、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里      巴巴(1688)、其他(OTHERS)
func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetChannel() string {
    return r._channel
}
// Trades Setter
// 订单信息，目前一次请求只支持一个主订单
func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetTrades(_trades []TradeOrderInfoDto) error {
    r._trades = _trades
    r.Set("trades", _trades)
    return nil
}

// Trades Getter
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetTrades() []TradeOrderInfoDto {
    return r._trades
}
// Receiver Setter
// 收件人信息
func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetReceiver(_receiver *ReceiverDto) error {
    r._receiver = _receiver
    r.Set("receiver", _receiver)
    return nil
}

// Receiver Getter
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetReceiver() *ReceiverDto {
    return r._receiver
}
// SdtExtendInfoDTO Setter
// 扩展信息
func (r *TaobaoOmniorderStoreGetconsignmailcodeRequest) SetSdtExtendInfoDTO(_sdtExtendInfoDTO *SdtExtendInfoDto) error {
    r._sdtExtendInfoDTO = _sdtExtendInfoDTO
    r.Set("sdt_extend_info_d_t_o", _sdtExtendInfoDTO)
    return nil
}

// SdtExtendInfoDTO Getter
func (r TaobaoOmniorderStoreGetconsignmailcodeRequest) GetSdtExtendInfoDTO() *SdtExtendInfoDto {
    return r._sdtExtendInfoDTO
}