package tblogistics

import (
	"net/url"
	"sync"

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
	// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。instantMobilePhoneNumber表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号"12345678910-1234"。
	_feature string
	// 淘宝交易ID
	_tid string
	// 发货的子订单id列表（consign_type = 1、2、3 时不再使用次字段，使用新字段goods代替需要发货的子订单信息）
	_subTid string
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
	_cancelId int64
	// 子订单发货状态
	_consignStatus *ConsignStatusRequest
	// 发货类型 0：普通发货(老链路) 1: 普通发货（新链路，支持子订单部分发货、成分品发货以及ERP线下赠品发货） 2: 将发货状态从"部分发"修改为"全部发" 3：补发；默认为0
	_consignType int64
}

// NewAlibabaAscpLogisticsOfflineSendRequest 初始化AlibabaAscpLogisticsOfflineSendAPIRequest对象
func NewAlibabaAscpLogisticsOfflineSendRequest() *AlibabaAscpLogisticsOfflineSendAPIRequest {
	return &AlibabaAscpLogisticsOfflineSendAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) Reset() {
	r._consignPkgs = r._consignPkgs[:0]
	r._feature = ""
	r._tid = ""
	r._subTid = ""
	r._senderId = 0
	r._cancelId = 0
	r._consignStatus = nil
	r._consignType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.offline.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetRawParams() model.Params {
	return r.Params
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
// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 &#34;|&#34;不同商品间的分隔符。 例1商品和2商品，之间就用&#34;|&#34;分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) &#34;:&#34;TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 &#34;,&#34; 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为&#34;12345&#34;,&#34;67890&#34;。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的&#34;12345&#34;,&#34;67890&#34;.说明本订单A宝贝的数量为2，值分别为&#34;12345&#34;,&#34;67890&#34;。 当存在&#34;|&#34;时，就说明订单中存在多个商品，商品间用&#34;|&#34;分隔了开来。|&#34;之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。instantMobilePhoneNumber表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号&#34;12345678910-1234&#34;。
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
// 发货的子订单id列表（consign_type = 1、2、3 时不再使用次字段，使用新字段goods代替需要发货的子订单信息）
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

// SetConsignStatus is ConsignStatus Setter
// 子订单发货状态
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetConsignStatus(_consignStatus *ConsignStatusRequest) error {
	r._consignStatus = _consignStatus
	r.Set("consign_status", _consignStatus)
	return nil
}

// GetConsignStatus ConsignStatus Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetConsignStatus() *ConsignStatusRequest {
	return r._consignStatus
}

// SetConsignType is ConsignType Setter
// 发货类型 0：普通发货(老链路) 1: 普通发货（新链路，支持子订单部分发货、成分品发货以及ERP线下赠品发货） 2: 将发货状态从&#34;部分发&#34;修改为&#34;全部发&#34; 3：补发；默认为0
func (r *AlibabaAscpLogisticsOfflineSendAPIRequest) SetConsignType(_consignType int64) error {
	r._consignType = _consignType
	r.Set("consign_type", _consignType)
	return nil
}

// GetConsignType ConsignType Getter
func (r AlibabaAscpLogisticsOfflineSendAPIRequest) GetConsignType() int64 {
	return r._consignType
}

var poolAlibabaAscpLogisticsOfflineSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsOfflineSendRequest()
	},
}

// GetAlibabaAscpLogisticsOfflineSendRequest 从 sync.Pool 获取 AlibabaAscpLogisticsOfflineSendAPIRequest
func GetAlibabaAscpLogisticsOfflineSendAPIRequest() *AlibabaAscpLogisticsOfflineSendAPIRequest {
	return poolAlibabaAscpLogisticsOfflineSendAPIRequest.Get().(*AlibabaAscpLogisticsOfflineSendAPIRequest)
}

// ReleaseAlibabaAscpLogisticsOfflineSendAPIRequest 将 AlibabaAscpLogisticsOfflineSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsOfflineSendAPIRequest(v *AlibabaAscpLogisticsOfflineSendAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsOfflineSendAPIRequest.Put(v)
}
