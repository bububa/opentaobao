package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaImapCategoryPredictAPIRequest 类目预测接口 API请求
// alibaba.imap.category.predict
//
// * 类目预测接口
//      * 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
//      * 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
type AlibabaImapCategoryPredictAPIRequest struct {
	model.Params
	// 入参DO
	_topImapItemDo *TopImapItemDo
	// 账号信息
	_fixedMappingAppInfo *FixedMappingAppInfo
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
func (r AlibabaImapCategoryPredictAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TopImapItemDo Setter
// 入参DO
func (r *AlibabaImapCategoryPredictAPIRequest) SetTopImapItemDo(_topImapItemDo *TopImapItemDo) error {
	r._topImapItemDo = _topImapItemDo
	r.Set("top_imap_item_do", _topImapItemDo)
	return nil
}

// Get TopImapItemDo Getter
func (r AlibabaImapCategoryPredictAPIRequest) GetTopImapItemDo() *TopImapItemDo {
	return r._topImapItemDo
}

// Set is FixedMappingAppInfo Setter
// 账号信息
func (r *AlibabaImapCategoryPredictAPIRequest) SetFixedMappingAppInfo(_fixedMappingAppInfo *FixedMappingAppInfo) error {
	r._fixedMappingAppInfo = _fixedMappingAppInfo
	r.Set("fixed_mapping_app_info", _fixedMappingAppInfo)
	return nil
}

// Get FixedMappingAppInfo Getter
func (r AlibabaImapCategoryPredictAPIRequest) GetFixedMappingAppInfo() *FixedMappingAppInfo {
	return r._fixedMappingAppInfo
}
