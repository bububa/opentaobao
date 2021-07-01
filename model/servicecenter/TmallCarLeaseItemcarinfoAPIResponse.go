package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseItemcarinfoAPIResponse
整车租赁商品四级车型信息 API返回值
tmall.car.lease.itemcarinfo

整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传 */
type TmallCarLeaseItemcarinfoAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseItemcarinfoAPIResponseModel
}

// TmallCarLeaseItemcarinfoAPIResponseModel is 整车租赁商品四级车型信息 成功返回结果
type TmallCarLeaseItemcarinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_itemcarinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallCarLeaseItemcarinfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
