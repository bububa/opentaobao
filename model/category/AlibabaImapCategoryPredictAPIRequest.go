package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaImapCategoryPredictAPIRequest
类目预测接口 API请求
alibaba.imap.category.predict

* 类目预测接口
     * 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
     * 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList */
type AlibabaImapCategoryPredictAPIRequest struct {
	model.Params
	// 入参DO
	_topImapItemDo *TopImapItemDo
	// 账号信息
	_fixedMappingAppInfo *FixedMappingAppInfo
}

// New
