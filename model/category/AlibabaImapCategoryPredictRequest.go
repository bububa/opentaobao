package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
类目预测接口 APIRequest
alibaba.imap.category.predict

* 类目预测接口
     * 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
     * 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
*/
type AlibabaImapCategoryPredictRequest struct {
    model.Params

    // 入参DO
    topImapItemDo   *TopImapItemDo 

    // 账号信息
    fixedMappingAppInfo   *FixedMappingAppInfo 

}

func NewAlibabaImapCategoryPredictRequest() *AlibabaImapCategoryPredictRequest{
    return &AlibabaImapCategoryPredictRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaImapCategoryPredictRequest) GetApiMethodName() string {
    return "alibaba.imap.category.predict"
}

func (r AlibabaImapCategoryPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaImapCategoryPredictRequest) SetTopImapItemDo(topImapItemDo *TopImapItemDo) error {
    r.topImapItemDo = topImapItemDo
    r.Set("top_imap_item_do", topImapItemDo)
    return nil
}

func (r AlibabaImapCategoryPredictRequest) GetTopImapItemDo() *TopImapItemDo {
    return r.topImapItemDo
}

func (r *AlibabaImapCategoryPredictRequest) SetFixedMappingAppInfo(fixedMappingAppInfo *FixedMappingAppInfo) error {
    r.fixedMappingAppInfo = fixedMappingAppInfo
    r.Set("fixed_mapping_app_info", fixedMappingAppInfo)
    return nil
}

func (r AlibabaImapCategoryPredictRequest) GetFixedMappingAppInfo() *FixedMappingAppInfo {
    return r.fixedMappingAppInfo
}

