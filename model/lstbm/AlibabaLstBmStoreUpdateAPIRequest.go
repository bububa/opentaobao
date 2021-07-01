package lstbm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstBmStoreUpdateAPIRequest
修改品牌商自有门店数据 API请求
alibaba.lst.bm.store.update

修改品牌商自有门店数据 */
type AlibabaLstBmStoreUpdateAPIRequest struct {
	model.Params
	// 门店数据模型
	_openStoreDto *LstTopOpenStoreDto
}

// New
