package lstbm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstBmStoreAddAPIResponse
导入品牌商自有门店 API返回值
alibaba.lst.bm.store.add

导入品牌商自有门店 */
type AlibabaLstBmStoreAddAPIResponse struct {
	model.CommonResponse
	AlibabaLstBmStoreAddAPIResponseModel
}

// AlibabaLstBmStoreAddAPIResponseModel is 导入品牌商自有门店 成功返回结果
type AlibabaLstBmStoreAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_bm_store_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示执行成功，false表示执行失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
