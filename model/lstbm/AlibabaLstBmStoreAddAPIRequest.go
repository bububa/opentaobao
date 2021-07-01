package lstbm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstBmStoreAddAPIRequest
导入品牌商自有门店 API请求
alibaba.lst.bm.store.add

导入品牌商自有门店 */
type AlibabaLstBmStoreAddAPIRequest struct {
	model.Params
	// 门店数据模型
	_openStoreDto *LstTopOpenStoreDto
}

// New
