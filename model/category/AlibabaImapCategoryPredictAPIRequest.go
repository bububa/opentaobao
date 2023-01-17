package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaImapCategoryPredictAPIRequest 类目预测接口 API请求
// alibaba.imap.category.predict
//
// * 类目预测接口
//   - 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
//   - 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
type AlibabaImapCategoryPredictAPIRequest struct {
	model.Params
	// 账号信息
	_fixedMappingAppInfo *FixedMappingAppInfo
	// 入参DO
	_topImapItemDo *TopImapItemDo
}

// NewAlibabaImapCategoryPredictRequest 初始化AlibabaImapCategoryPredictAPIRequest对象
func NewAlibabaImapCategoryPredictRequest() *AlibabaImapCategoryPredictAPIRequest {
	return &AlibabaImapCategoryPredictAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaImapCategoryPredictAPIRequest) GetApiMethodName() string {
	return "alibaba.imap.category.predict"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaImapCategoryPredictAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaImapCategoryPredictAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFixedMappingAppInfo is FixedMappingAppInfo Setter
// 账号信息
func (r *AlibabaImapCategoryPredictAPIRequest) SetFixedMappingAppInfo(_fixedMappingAppInfo *FixedMappingAppInfo) error {
	r._fixedMappingAppInfo = _fixedMappingAppInfo
	r.Set("fixed_mapping_app_info", _fixedMappingAppInfo)
	return nil
}

// GetFixedMappingAppInfo FixedMappingAppInfo Getter
func (r AlibabaImapCategoryPredictAPIRequest) GetFixedMappingAppInfo() *FixedMappingAppInfo {
	return r._fixedMappingAppInfo
}

// SetTopImapItemDo is TopImapItemDo Setter
// 入参DO
func (r *AlibabaImapCategoryPredictAPIRequest) SetTopImapItemDo(_topImapItemDo *TopImapItemDo) error {
	r._topImapItemDo = _topImapItemDo
	r.Set("top_imap_item_do", _topImapItemDo)
	return nil
}

// GetTopImapItemDo TopImapItemDo Getter
func (r AlibabaImapCategoryPredictAPIRequest) GetTopImapItemDo() *TopImapItemDo {
	return r._topImapItemDo
}
