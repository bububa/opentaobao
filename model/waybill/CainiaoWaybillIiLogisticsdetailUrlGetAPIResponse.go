package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiLogisticsdetailUrlGetAPIResponse
电子面单物流详情授权url获取 API返回值
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url */
type CainiaoWaybillIiLogisticsdetailUrlGetAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiLogisticsdetailUrlGetAPIResponseModel
}

// CainiaoWaybillIiLogisticsdetailUrlGetAPIResponseModel is 电子面单物流详情授权url获取 成功返回结果
type CainiaoWaybillIiLogisticsdetailUrlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_logisticsdetail_url_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权访问的url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
