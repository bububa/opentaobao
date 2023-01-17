package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest 点击标签后排序接口 API请求
// alibaba.alihealth.pregnancy.taginfo.addrank
//
// 备孕管理--点击标签后排序接口
type AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest struct {
	model.Params
	// 标签编码，例如备孕标签为5122
	_tagCode string
	// 用户id
	_userId int64
}

// NewAlibabaAlihealthPregnancyTaginfoAddrankRequest 初始化AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest对象
func NewAlibabaAlihealthPregnancyTaginfoAddrankRequest() *AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest {
	return &AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.taginfo.addrank"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagCode is TagCode Setter
// 标签编码，例如备孕标签为5122
func (r *AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest) SetTagCode(_tagCode string) error {
	r._tagCode = _tagCode
	r.Set("tag_code", _tagCode)
	return nil
}

// GetTagCode TagCode Getter
func (r AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest) GetTagCode() string {
	return r._tagCode
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest) GetUserId() int64 {
	return r._userId
}
