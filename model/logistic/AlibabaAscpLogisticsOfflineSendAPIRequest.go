package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsOfflineSendAPIRequest 自己联系物流发货 API请求
// alibaba.ascp.logistics.offline.send
//
// 用户调用该接口可实现自己联系发货，使用该接口发货，交易订单状态会直接变成卖家已发货
type AlibabaAscpLogisticsOfflineSendAPIRequest struct {
	model.Params
	// 包裹信息
	_consignPkgs []TopConsignPkgRequest
	// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段
	_feature string
	// 淘宝交易ID
	_tid string
	// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
	_subTid string
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
	_cancelId int64
}

// NewAlibabaAscpLogisticsOfflineSendRequest 初始化AlibabaAscpLogisticsOfflineSendAPIRequest对象
func NewAlibabaAscpLogisticsOfflineSendRequest() *AlibabaAscpLogisticsOfflineSendAPIRequest {
	return &AlibabaAscpLogisticsOfflineSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.offline.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetConsignPkgs is ConsignPkgs Setter
// 包裹信息
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetConsignPkgs(_consignPkgs []TopConsignPkgRequest) error {
	r._consignPkgs = _consignPkgs
	r.Set("consign_pkgs", _consignPkgs)
	return nil
}

// GetConsignPkgs ConsignPkgs Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetConsignPkgs() []TopConsignPkgRequest {
	return r._consignPkgs
}

// SetFeature is Feature Setter
// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetFeature() string {
	return r._feature
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetTid(_tid string) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetTid() string {
	return r._tid
}

// SetSubTid is SubTid Setter
// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetSubTid(_subTid string) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// GetSubTid SubTid Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetSubTid() string {
	return r._subTid
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetCancelId is CancelId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetCancelId(_cancelId int64) error {
	r._cancelId = _cancelId
	r.Set("cancel_id", _cancelId)
	return nil
}

// GetCancelId CancelId Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetCancelId() int64 {
	return r._cancelId
}
