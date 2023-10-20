package lstpos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopeninventorysyncinventorydataAPIResponse 商品库存修改同步接口(最多20条库存信息) API返回值
// alibaba.lst.pos.open.inventory.syncinventorydata
//
// 商品库存修改同步接口(最多20条库存信息)
type AlibabalstposopeninventorysyncinventorydataAPIResponse struct {
	model.CommonResponse
	AlibabalstposopeninventorysyncinventorydataAPIResponseModel
}

// AlibabalstposopeninventorysyncinventorydataAPIResponseModel is 商品库存修改同步接口(最多20条库存信息) 成功返回结果
type AlibabalstposopeninventorysyncinventorydataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_inventory_syncinventorydata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlibabalstposopeninventorysyncinventorydataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
