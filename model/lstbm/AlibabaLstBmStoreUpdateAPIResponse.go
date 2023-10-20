package lstbm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstbmstoreupdateAPIResponse 修改品牌商自有门店数据 API返回值
// alibaba.lst.bm.store.update
//
// 修改品牌商自有门店数据
type AlibabalstbmstoreupdateAPIResponse struct {
	model.CommonResponse
	AlibabalstbmstoreupdateAPIResponseModel
}

// AlibabalstbmstoreupdateAPIResponseModel is 修改品牌商自有门店数据 成功返回结果
type AlibabalstbmstoreupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_bm_store_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示执行成功，false表示执行失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
