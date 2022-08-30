package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketStoreGetAPIResponse 查询电子凭证对应门店信息 API返回值
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
type TmallAliautoEticketStoreGetAPIResponse struct {
	model.CommonResponse
	TmallAliautoEticketStoreGetAPIResponseModel
}

// TmallAliautoEticketStoreGetAPIResponseModel is 查询电子凭证对应门店信息 成功返回结果
type TmallAliautoEticketStoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_eticket_store_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
