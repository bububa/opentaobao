package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsSellerSendAPIRequest 商家配送发货 API请求
// alibaba.ascp.logistics.seller.send
//
// 该API提供商家配送发货能力，使用该接口发货，交易订单状态会直接变成卖家已发货
type AlibabaAscpLogisticsSellerSendAPIRequest struct {
	model.Params
	// 派送员手机号（支持座机号和400）
	_deliveryMobile string
	// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段
	_feature string
	// 淘宝交易ID
	_tid string
	// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
	_subTid string
	// 派送员姓名
	_deliveryName string
	//  A（默认，核销签收模式），B（商家回传物流节点模式）
	_mode string
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
	_cancelId int64
}

// NewAlibabaAscpLogisticsSellerSendRequest 初始化AlibabaAscpLogisticsSellerSendAPIRequest对象
func NewAlibabaAscpLogisticsSellerSendRequest() *AlibabaAscpLogisticsSellerSendAPIRequest {
	return &AlibabaAscpLogisticsSellerSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.seller.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryMobile is DeliveryMobile Setter
// 派送员手机号（支持座机号和400）
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetDeliveryMobile(_deliveryMobile string) error {
	r._deliveryMobile = _deliveryMobile
	r.Set("delivery_mobile", _deliveryMobile)
	return nil
}

// GetDeliveryMobile DeliveryMobile Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetDeliveryMobile() string {
	return r._deliveryMobile
}

// SetFeature is Feature Setter
// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 &#34;|&#34;不同商品间的分隔符。 例1商品和2商品，之间就用&#34;|&#34;分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) &#34;:&#34;TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 &#34;,&#34; 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为&#34;12345&#34;,&#34;67890&#34;。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的&#34;12345&#34;,&#34;67890&#34;.说明本订单A宝贝的数量为2，值分别为&#34;12345&#34;,&#34;67890&#34;。 当存在&#34;|&#34;时，就说明订单中存在多个商品，商品间用&#34;|&#34;分隔了开来。|&#34;之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetFeature() string {
	return r._feature
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetTid(_tid string) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetTid() string {
	return r._tid
}

// SetSubTid is SubTid Setter
// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetSubTid(_subTid string) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// GetSubTid SubTid Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetSubTid() string {
	return r._subTid
}

// SetDeliveryName is DeliveryName Setter
// 派送员姓名
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetDeliveryName(_deliveryName string) error {
	r._deliveryName = _deliveryName
	r.Set("delivery_name", _deliveryName)
	return nil
}

// GetDeliveryName DeliveryName Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetDeliveryName() string {
	return r._deliveryName
}

// SetMode is Mode Setter
//
//	A（默认，核销签收模式），B（商家回传物流节点模式）
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetMode(_mode string) error {
	r._mode = _mode
	r.Set("mode", _mode)
	return nil
}

// GetMode Mode Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetMode() string {
	return r._mode
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetCancelId is CancelId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
func (r *AlibabaAscpLogisticsSellerSendAPIRequest) SetCancelId(_cancelId int64) error {
	r._cancelId = _cancelId
	r.Set("cancel_id", _cancelId)
	return nil
}

// GetCancelId CancelId Getter
func (r AlibabaAscpLogisticsSellerSendAPIRequest) GetCancelId() int64 {
	return r._cancelId
}
