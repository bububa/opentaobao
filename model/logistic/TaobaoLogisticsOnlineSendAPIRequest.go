package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOnlineSendAPIRequest 在线订单发货处理（支持货到付款） API请求
// taobao.logistics.online.send
//
// 用户调用该接口可实现在线订单发货（支持货到付款）<br/>调用该接口实现在线下单发货，有两种情况：<br><br/><font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br><br/>如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font>
type TaobaoLogisticsOnlineSendAPIRequest struct {
	model.Params
	// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
	_subTid []int64
	// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
	_outSid string
	// 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
	_companyCode string
	// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
	_feature string
	// 商家的IP地址
	_sellerIp string
	// 淘宝交易ID
	_tid int64
	// 表明是否是拆单，默认值0，1表示拆单
	_isSplit int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
	_cancelId int64
}

// NewTaobaoLogisticsOnlineSendRequest 初始化TaobaoLogisticsOnlineSendAPIRequest对象
func NewTaobaoLogisticsOnlineSendRequest() *TaobaoLogisticsOnlineSendAPIRequest {
	return &TaobaoLogisticsOnlineSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOnlineSendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.online.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOnlineSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubTid is SubTid Setter
// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetSubTid(_subTid []int64) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// GetSubTid SubTid Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetSubTid() []int64 {
	return r._subTid
}

// SetOutSid is OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetOutSid(_outSid string) error {
	r._outSid = _outSid
	r.Set("out_sid", _outSid)
	return nil
}

// GetOutSid OutSid Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetOutSid() string {
	return r._outSid
}

// SetCompanyCode is CompanyCode Setter
// 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetCompanyCode(_companyCode string) error {
	r._companyCode = _companyCode
	r.Set("company_code", _companyCode)
	return nil
}

// GetCompanyCode CompanyCode Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetCompanyCode() string {
	return r._companyCode
}

// SetFeature is Feature Setter
// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetFeature() string {
	return r._feature
}

// SetSellerIp is SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetSellerIp(_sellerIp string) error {
	r._sellerIp = _sellerIp
	r.Set("seller_ip", _sellerIp)
	return nil
}

// GetSellerIp SellerIp Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetSellerIp() string {
	return r._sellerIp
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetTid() int64 {
	return r._tid
}

// SetIsSplit is IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetIsSplit(_isSplit int64) error {
	r._isSplit = _isSplit
	r.Set("is_split", _isSplit)
	return nil
}

// GetIsSplit IsSplit Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetIsSplit() int64 {
	return r._isSplit
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetCancelId is CancelId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
func (r *TaobaoLogisticsOnlineSendAPIRequest) SetCancelId(_cancelId int64) error {
	r._cancelId = _cancelId
	r.Set("cancel_id", _cancelId)
	return nil
}

// GetCancelId CancelId Getter
func (r TaobaoLogisticsOnlineSendAPIRequest) GetCancelId() int64 {
	return r._cancelId
}
