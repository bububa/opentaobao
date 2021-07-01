package lstbm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstBmStoreEmpSaveAPIRequest
保存品牌商自有门店和内部业代之间的关系 API请求
alibaba.lst.bm.store.emp.save

保存品牌商自有门店和内部业代之间的关系 */
type AlibabaLstBmStoreEmpSaveAPIRequest struct {
	model.Params
	// 门店id标识
	_storeId string
	// 员工id标识
	_bmEmpId string
}

// New
