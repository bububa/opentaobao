package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthLabItemStoreRelationSyncAPIRequest
检验检测业务，isv项目门店关系同步 API请求
alibaba.alihealth.lab.item.store.relation.sync

阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除 */
type AlibabaAlihealthLabItemStoreRelationSyncAPIRequest struct {
	model.Params
	// EFFECTIVE 有效，INVALID 无效
	_isvRelationStatus string
	// isv门店编码
	_isvStoreCodes []string
	// 检验检测项目isv侧编码
	_isvItemCode string
}

// New
