package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
类目预测接口 API请求
alibaba.imap.category.predict

* 类目预测接口
     * 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
     * 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
*/
type AlibabaImapCategoryPredictRequest struct {
    model.Params
    // 入参DO
    _topImapItemDo   *TopImapItemDo
    // 账号信息
    _fixedMappingAppInfo   *FixedMappingAppInfo
}

// 初始化AlibabaImapCategoryPredictRequest对象
func NewAlibabaImapCategoryPredictRequest() *AlibabaImapCategoryPredictRequest{
    return &AlibabaImapCategoryPredictRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaImapCategoryPredictRequest) GetApiMethodName() string {
    return "alibaba.imap.category.predict"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaImapCategoryPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopImapItemDo Setter
// 入参DO
func (r *AlibabaImapCategoryPredictRequest) SetTopImapItemDo(_topImapItemDo *TopImapItemDo) error {
    r._topImapItemDo = _topImapItemDo
    r.Set("top_imap_item_do", _topImapItemDo)
    return nil
}

// TopImapItemDo Getter
func (r AlibabaImapCategoryPredictRequest) GetTopImapItemDo() *TopImapItemDo {
    return r._topImapItemDo
}
// FixedMappingAppInfo Setter
// 账号信息
func (r *AlibabaImapCategoryPredictRequest) SetFixedMappingAppInfo(_fixedMappingAppInfo *FixedMappingAppInfo) error {
    r._fixedMappingAppInfo = _fixedMappingAppInfo
    r.Set("fixed_mapping_app_info", _fixedMappingAppInfo)
    return nil
}

// FixedMappingAppInfo Getter
func (r AlibabaImapCategoryPredictRequest) GetFixedMappingAppInfo() *FixedMappingAppInfo {
    return r._fixedMappingAppInfo
}
