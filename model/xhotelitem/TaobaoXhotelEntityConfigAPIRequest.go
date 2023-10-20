package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelentityconfigAPIRequest 飞猪商品各实体通用配置 API请求
// taobao.xhotel.entity.config
//
// 飞猪商品各实体通用配置服务
type TaobaoxhotelentityconfigAPIRequest struct {
	model.Params
	// 实体编码，例如酒店编码，房价编码
	_entityCode string
	// 格式为：[{"invoice_status":"1"},{"invoice_has":"1"},{"invoice_canvat":"1"},{"invoice_provider":"2"},{"invoice_channel":"1"},{"invoice_payer":"1"},{"invoice_type":"1,2,3"},{"invoice_desc":"测试发票描叙123"}]  key-value形式的配置信息其中invoice_status为发票状态，取值为0(无效),1(有效);invoice_has是否能开发票，1(能),0(不能);invoice_canvat是否能提供增值税专用发票1(能),0(不能)；invoice_provider发票提供方1(酒店提供),2(卖家邮寄)；invoice_channel开票渠道，目前仅支持1(线下开具)；invoice_payer邮费承担方，1(卖家承担)，2(买家承担-邮费到付);invoice_type开票类型 1(代订房费),2(旅游费) 3(差旅费) 4(卖家提供Email版收据)  如果多个请以英文逗号隔开，例如1,2,3。invoice_desc发票补充说明，如果没有可以填写无，长度不要超过20字符。该接口是全量更新的形式，每次更新必须要携带这些节点信息，如果想失效掉该rp上的发票信息，那么可以将invoice_status设置为0。
	_configData string
	// 默认是taobao。和酒店，房型等实体上的vendor保持一致。
	_vendor string
	// 实体id,例如酒店id,房型id(飞猪平台的id),rpid等，目前仅支持rpid
	_entityId int64
	// 实体类型：1：卖家维度，2：酒店维度，3：房型维度，4：宝贝维度，5：rp维度。请注意目前仅支持类型5
	_type int64
}

// NewTaobaoxhotelentityconfigRequest 初始化TaobaoxhotelentityconfigAPIRequest对象
func NewTaobaoxhotelentityconfigRequest() *TaobaoxhotelentityconfigAPIRequest {
	return &TaobaoxhotelentityconfigAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelentityconfigAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.entity.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelentityconfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelentityconfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntityCode is EntityCode Setter
// 实体编码，例如酒店编码，房价编码
func (r *TaobaoxhotelentityconfigAPIRequest) SetEntityCode(_entityCode string) error {
	r._entityCode = _entityCode
	r.Set("entity_code", _entityCode)
	return nil
}

// GetEntityCode EntityCode Getter
func (r TaobaoxhotelentityconfigAPIRequest) GetEntityCode() string {
	return r._entityCode
}

// SetConfigData is ConfigData Setter
// 格式为：[{&#34;invoice_status&#34;:&#34;1&#34;},{&#34;invoice_has&#34;:&#34;1&#34;},{&#34;invoice_canvat&#34;:&#34;1&#34;},{&#34;invoice_provider&#34;:&#34;2&#34;},{&#34;invoice_channel&#34;:&#34;1&#34;},{&#34;invoice_payer&#34;:&#34;1&#34;},{&#34;invoice_type&#34;:&#34;1,2,3&#34;},{&#34;invoice_desc&#34;:&#34;测试发票描叙123&#34;}]  key-value形式的配置信息其中invoice_status为发票状态，取值为0(无效),1(有效);invoice_has是否能开发票，1(能),0(不能);invoice_canvat是否能提供增值税专用发票1(能),0(不能)；invoice_provider发票提供方1(酒店提供),2(卖家邮寄)；invoice_channel开票渠道，目前仅支持1(线下开具)；invoice_payer邮费承担方，1(卖家承担)，2(买家承担-邮费到付);invoice_type开票类型 1(代订房费),2(旅游费) 3(差旅费) 4(卖家提供Email版收据)  如果多个请以英文逗号隔开，例如1,2,3。invoice_desc发票补充说明，如果没有可以填写无，长度不要超过20字符。该接口是全量更新的形式，每次更新必须要携带这些节点信息，如果想失效掉该rp上的发票信息，那么可以将invoice_status设置为0。
func (r *TaobaoxhotelentityconfigAPIRequest) SetConfigData(_configData string) error {
	r._configData = _configData
	r.Set("config_data", _configData)
	return nil
}

// GetConfigData ConfigData Getter
func (r TaobaoxhotelentityconfigAPIRequest) GetConfigData() string {
	return r._configData
}

// SetVendor is Vendor Setter
// 默认是taobao。和酒店，房型等实体上的vendor保持一致。
func (r *TaobaoxhotelentityconfigAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelentityconfigAPIRequest) GetVendor() string {
	return r._vendor
}

// SetEntityId is EntityId Setter
// 实体id,例如酒店id,房型id(飞猪平台的id),rpid等，目前仅支持rpid
func (r *TaobaoxhotelentityconfigAPIRequest) SetEntityId(_entityId int64) error {
	r._entityId = _entityId
	r.Set("entity_id", _entityId)
	return nil
}

// GetEntityId EntityId Getter
func (r TaobaoxhotelentityconfigAPIRequest) GetEntityId() int64 {
	return r._entityId
}

// SetType is Type Setter
// 实体类型：1：卖家维度，2：酒店维度，3：房型维度，4：宝贝维度，5：rp维度。请注意目前仅支持类型5
func (r *TaobaoxhotelentityconfigAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoxhotelentityconfigAPIRequest) GetType() int64 {
	return r._type
}
