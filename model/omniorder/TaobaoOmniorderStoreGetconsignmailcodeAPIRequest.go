package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreGetconsignmailcodeAPIRequest
全渠道门店物流菜鸟裹裹取号 API请求
taobao.omniorder.store.getconsignmailcode

用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号 */
type TaobaoOmniorderStoreGetconsignmailcodeAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 发件人联系电话，如空则表示使用门店信息中的电话号码
	_senderContact string
	// 淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)      、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)      、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里      巴巴(1688)、其他(OTHERS)
	_channel string
	// 订单信息，目前一次请求只支持一个主订单
	_trades []TradeOrderInfoDto
	// 收件人信息
	_receiver *ReceiverDto
	// 扩展信息
	_sdtExtendInfoDTO *SdtExtendInfoDto
}

// New
