package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsConsignModifyAPIRequest 修改物流公司和运单号 API请求
// alibaba.ascp.logistics.consign.modify
//
// 修改物流公司和运单号
type AlibabaAscpLogisticsConsignModifyAPIRequest struct {
	model.Params
	// 订单id
	_tid string
	// 原物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取
	_oldCompanyCode string
	// 原运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入
	_oldOutSid string
	// 新物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取
	_newCompanyCode string
	// 新运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入
	_newOutSid string
	// feature参数格式，KV之间用“=”分隔，多个key之间用”;”分隔 ，范例: instantMobilePhoneNumber=12345678910表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号"12345678910-1234"。
	_feature string
	// 原包裹中的商品信息
	_goods *TopConsignGoodsRequest
}

// NewAlibabaAscpLogisticsConsignModifyRequest 初始化AlibabaAscpLogisticsConsignModifyAPIRequest对象
func NewAlibabaAscpLogisticsConsignModifyRequest() *AlibabaAscpLogisticsConsignModifyAPIRequest {
	return &AlibabaAscpLogisticsConsignModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.consign.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 订单id
func (r *AlibabaAscpLogisticsConsignModifyAPIRequest) SetTid(_tid string) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetTid() string {
	return r._tid
}

// SetOldCompanyCode is OldCompanyCode Setter
// 原物流公司代码.如&#34;POST&#34;代表中国邮政,&#34;ZJS&#34;代表宅急送。调用 taobao.logistics.companies.get 获取
func (r *AlibabaAscpLogisticsConsignModifyAPIRequest) SetOldCompanyCode(_oldCompanyCode string) error {
	r._oldCompanyCode = _oldCompanyCode
	r.Set("old_company_code", _oldCompanyCode)
	return nil
}

// GetOldCompanyCode OldCompanyCode Getter
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetOldCompanyCode() string {
	return r._oldCompanyCode
}

// SetOldOutSid is OldOutSid Setter
// 原运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入
func (r *AlibabaAscpLogisticsConsignModifyAPIRequest) SetOldOutSid(_oldOutSid string) error {
	r._oldOutSid = _oldOutSid
	r.Set("old_out_sid", _oldOutSid)
	return nil
}

// GetOldOutSid OldOutSid Getter
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetOldOutSid() string {
	return r._oldOutSid
}

// SetNewCompanyCode is NewCompanyCode Setter
// 新物流公司代码.如&#34;POST&#34;代表中国邮政,&#34;ZJS&#34;代表宅急送。调用 taobao.logistics.companies.get 获取
func (r *AlibabaAscpLogisticsConsignModifyAPIRequest) SetNewCompanyCode(_newCompanyCode string) error {
	r._newCompanyCode = _newCompanyCode
	r.Set("new_company_code", _newCompanyCode)
	return nil
}

// GetNewCompanyCode NewCompanyCode Getter
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetNewCompanyCode() string {
	return r._newCompanyCode
}

// SetNewOutSid is NewOutSid Setter
// 新运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入
func (r *AlibabaAscpLogisticsConsignModifyAPIRequest) SetNewOutSid(_newOutSid string) error {
	r._newOutSid = _newOutSid
	r.Set("new_out_sid", _newOutSid)
	return nil
}

// GetNewOutSid NewOutSid Getter
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetNewOutSid() string {
	return r._newOutSid
}

// SetFeature is Feature Setter
// feature参数格式，KV之间用“=”分隔，多个key之间用”;”分隔 ，范例: instantMobilePhoneNumber=12345678910表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号&#34;12345678910-1234&#34;。
func (r *AlibabaAscpLogisticsConsignModifyAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetFeature() string {
	return r._feature
}

// SetGoods is Goods Setter
// 原包裹中的商品信息
func (r *AlibabaAscpLogisticsConsignModifyAPIRequest) SetGoods(_goods *TopConsignGoodsRequest) error {
	r._goods = _goods
	r.Set("goods", _goods)
	return nil
}

// GetGoods Goods Getter
func (r AlibabaAscpLogisticsConsignModifyAPIRequest) GetGoods() *TopConsignGoodsRequest {
	return r._goods
}
