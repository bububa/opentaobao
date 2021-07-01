package lstbm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstBmStoreEmpSaveAPIResponse
保存品牌商自有门店和内部业代之间的关系 API返回值
alibaba.lst.bm.store.emp.save

保存品牌商自有门店和内部业代之间的关系 */
type AlibabaLstBmStoreEmpSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLstBmStoreEmpSaveAPIResponseModel
}

// AlibabaLstBmStoreEmpSaveAPIResponseModel is 保存品牌商自有门店和内部业代之间的关系 成功返回结果
type AlibabaLstBmStoreEmpSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_bm_store_emp_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示执行成功，false表示执行失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
